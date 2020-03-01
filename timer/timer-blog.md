## Go-Timer源码解读

![](http://images.hcyhj.cn/blogimages/timer/timer-head.jpg)

### 前言

在初学Go定时任务之时，脑海中始终有一个问题在徘徊，究竟是每个任务都有一个goroutine去监控，还是多个任务处于同一个队列，让同一个goroutine去轮询检查。这里大家可以带着这个问题去进行接下来的阅读。

### Example

先来看一个简单的例子，这里我选择了`NewTicker`去进行测试，它和`NewTimer`唯一的区别是：前者定时循环执行，后者只会执行一次。

```go
func main() {
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			log.Println("xxxxxxxxxx")
		}
	}
}
```

这里会每隔5秒就会执行一次打印，但这个究竟是怎么实现的，咱们一步步的去探索。

### 源码部分



#### NewTicker

```go
func NewTicker(d Duration) *Ticker {
	if d <= 0 {
		panic(errors.New("non-positive interval for NewTicker"))
	}
	c := make(chan Time, 1)
	t := &Ticker{
		C: c,
		r: runtimeTimer{
			when:   when(d),
			period: int64(d),
			f:      sendTime,
			arg:    c,
		},
	}
	startTimer(&t.r)
	return t
}
```

简单的看一下第一个函数，不算太复杂，进行简单的异常判断，创建一个缓冲为1的`channel`，并构建核心的结构体`runtimeTimer`。下面我们关注一下这个结构体的几个属性。(由于`runtimeTimer`与`timer`底层结构一致，我这里截取`timer`结构体的源码进行解释一下相关属性)

```go
type timer struct {
	tb *timersBucket // 在哪个桶中存在，这里是根据goroutine所属的p确定的
	i  int           // 在堆结构中的索引位置
	when   int64 // 啥时候去执行函数f
	period int64 // 间隔多久去执行函数f，该值为0时表示只会执行一次函数f
	f      func(interface{}, uintptr)
	arg    interface{} // 函数f的第一个参数
	seq    uintptr // 函数f的第二个参数
}
```



#### startTimer

```go
func startTimer(*runtimeTimer)
```

接着看`startTimer`这个函数，初学者看这段源码时可能会觉得奇怪，因为它根本没有body。其实类似的情况并不少见，像这种没有方法体的大多都会在`runtime`包给其提供实现。如下所示：

```go
// startTimer adds t to the timer heap.
//go:linkname startTimer time.startTimer
func startTimer(t *timer) {
	if raceenabled {
		racerelease(unsafe.Pointer(t))
	}
	addtimer(t)
}
```

注意一下这个方法上面有一句注释`//go:linkname startTimer time.startTimer`，这句注释可不是一个无用的注释，**简单的来说`go:linkname`这个指令告诉编译器为当前源文件中私有函数或者变量在编译时链接到指定的方法或变量。**所以在这里大家可以把这个理解为`runtime.startTimer`是`time.startTimer`的具体实现。

#### addtimer

```go
func addtimer(t *timer) {
	tb := t.assignBucket()
	lock(&tb.lock)
	ok := tb.addtimerLocked(t)
	unlock(&tb.lock)
	if !ok {
		badTimer()
	}
}
```

这个函数主要干了两件事：

- 获取这个`timer`属于哪个`bucket`，这里是根据`goroutine`所属的`p`的id来进行计算。
- 将`timer`其添加到对应的`bucket`中。

这里可以保证在大部分情况下同一个`p`上创建的`timer`可以放到同一个`bucket`中，除非你的机器CPU核数超过了64个。每个核上维护着一个队列，在某种程度上也是提升了定时任务的性能。



#### addtimerLocked

```go
func (tb *timersBucket) addtimerLocked(t *timer) bool {
	// 保证when的值是正数
	if t.when < 0 {
		t.when = 1<<63 - 1
	}
	t.i = len(tb.t)
	tb.t = append(tb.t, t)
    // 根据when的值去调整堆中的顺序
	if !siftupTimer(tb.t, t.i) {
		return false
	}
	if t.i == 0 {
	
		if !tb.created {
			tb.created = true
            // 进行对当前bucket监控的goroutine的创建
			go timerproc(tb)
		}
	}
	return true
}
```

这里存储`timer`的数据结构是四叉树。相同的数据而言，四叉树比二叉树的深度要低，查询时效率要高一点。在实现定时器时为啥要选择四叉树而不是二叉树，大家可以参考一下这篇文章[定时器：4叉堆与2叉堆的效率比较](https://blog.csdn.net/znzxc/article/details/85916740)。

这个方法大概分三步：

- `tb.t = append(tb.t, t)`，将其插入到数组的最后一个。
- `siftupTimer(tb.t, t.i)`，将最后一个`timer`和它的parent进行比较，由于这里的数据结构是**四叉树**，所以它的parent计算公式为`p := (i - 1) / 4 // parent`，如果比parent小，则进行交换。这里会递归执行，直到取到符合`timer`的位置为止。
- 判断监控该`bucket`的goroutine是否已经创建，如果没有，则进行创建。

#### timerproc

```go
func timerproc(tb *timersBucket) {
	tb.gp = getg()
	for {
		lock(&tb.lock)
		tb.sleeping = false
		now := nanotime()
		delta := int64(-1)
		for {
			if len(tb.t) == 0 {
				delta = -1
				break
			}
            // 由于这里是最小堆，取出堆顶元素也就是最靠近执行时间的那个timer
			t := tb.t[0]
			delta = t.when - now
			if delta > 0 {
				break
			}
			ok := true
			if t.period > 0 {
                // 这里表示这个timer是一个定时轮询的任务，所以加上执行周期重新
                // 调整在堆中的位置
				// leave in heap but adjust next time to fire
				t.when += t.period * (1 + -delta/t.period)
                // 调整该timer在堆中的位置
				if !siftdownTimer(tb.t, 0) {
					ok = false
				}
			} else {
				// 只执行一次的任务，执行后直接从堆中移除
				last := len(tb.t) - 1
				if last > 0 {
					tb.t[0] = tb.t[last]
					tb.t[0].i = 0
				}
				tb.t[last] = nil
				tb.t = tb.t[:last]
				if last > 0 {
					if !siftdownTimer(tb.t, 0) {
						ok = false
					}
				}
				t.i = -1 // mark as removed
			}
			f := t.f
			arg := t.arg
			seq := t.seq
			unlock(&tb.lock)
			if !ok {
				badTimer()
			}
			if raceenabled {
				raceacquire(unsafe.Pointer(t))
			}
            // 执行timer结构中的f函数
			f(arg, seq)
			lock(&tb.lock)
		}
		if delta < 0 || faketime > 0 {
			// No timers left - put goroutine to sleep.
			tb.rescheduling = true
			goparkunlock(&tb.lock, waitReasonTimerGoroutineIdle, traceEvGoBlock, 1)
			continue
		}
		// At least one timer pending. Sleep until then.
		tb.sleeping = true
		tb.sleepUntil = now + delta
		noteclear(&tb.waitnote)
		unlock(&tb.lock)
		notetsleepg(&tb.waitnote, delta)
	}
}

```

执行流程：

1. 由于是最小堆，从堆顶取出的`timer`就是最近一个将要执行的任务，与当前时间进行对比，判断是否已经到了执行任务的时间。
2. 如果是定时轮询任务，取出来做好记录后需要调整该`timer`的属性`when`的值，并在堆中进行重新排序。方便下一次的执行。
3. 如果是执行一次的任务，取出来做好记录后需要从堆中进行移除。
4. 执行特定的函数，例如`sendTime`、`goFunc`等等函数。



### 总结

通过上面的了解咱们可以完美解决咱们在文章开始的时候提出的那个问题，究竟开了多少个goroutine去维护咱们的定时任务队列？答案是：比如你的机器有n个CPU，那么就会有n个`bucket`，同样就会有n个goroutine去监控这些`bucket`，由于存储结构采用的是最小堆，这里咱们也不用轮询检查，只用检查堆中的第一个元素即可。当然最后得出的结论并不属于咱们上面两个猜测的其中任何一个。