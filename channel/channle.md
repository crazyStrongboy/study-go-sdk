## channel

#### 简介

熟悉Go的人都知道，它提倡着不要通过共享内存来通讯，而要通过通讯来共享内存。Go提供了一种独特的并发同步技术来实现通过通讯来共享内存，此技术即为通道。我们可以把一个通道看作是在一个程序内部的一个FIFO数据队列。 一些协程可以向此通道发送数据，另外一些协程可以从此通道接收数据。



#### Example
介绍一下简单的用法：
```go
func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	t := <-c
	fmt.Println(t)
}
```
几个注意点：

1. 向一个nil通道中发送一个值，将会永久阻塞。
2. 向一个已关闭的通道中发送一个值，将会导致panic。
3. 可以从关闭的通道中读取值，缓冲区为空时，读取的是通道类型的零值。
4. 重复关闭一个通道也会导致panic。
5. 通道的元素值的传递都是复制过程，且至少被复制过一次以上。（直接复制到receiver中经过一次复制，通过缓冲区的话则经历了两次复制）



#### channel有两种类型，Unbuffered channels与Buffered channels

##### Unbuffered channels
```go
	c:=make(chan int)
```
它是一个阻塞型channel，必须要receiver也准备好的情况下，sender才能够将消息投递到c中去。可以结合下图进行思考一波：

![unbuffer_channel](images/unbuffer_channel.png)

##### Buffered channels
```go
	c:=make(chan int,1)
```
在buf未满之前，它是一个非阻塞型channel，sender可以将符合channel类型的值投递到channel中去，它内部会自己维护一个队列。当buf满了之后，sender会阻塞。可以结合下图进行思考一波：

![buffer_channel](images/buffer_channel.png)

#### 几种应用模式

##### for-range

```go
func forRange() {
	c := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		c <- i
	}
	close(c)
	for v := range c {
		fmt.Println(v)
	}
}
```

1. 在进行for-range一个通道时，该循环将源源不断的从通道中获取数据，直到此通道关闭并且它的缓冲队列中为空为止。
2. 这里的通道一定不能是单向发送通道（`chan <- int`）。
3. 当for-range一个空通道时，将会永久阻塞

##### select-case

```go
func selectCase() {
	c := make(chan int, 1)
	c <- 1
	close(c)
	select {
	case <-c:
		fmt.Println("xxxx")
	default:
		fmt.Println("aaaa")
	}
}
```

1. 每个`case`关键字后必须跟随一个通道接收数据操作或者一个通道发送数据操作。

2. 所有的非阻塞`case`操作中将有一个被随机选择执行（而不是按照从上到下的顺序），然后执行此操作对应的`case`分支代码块。

3. 在所有的`case`操作均为阻塞的情况下，如果`default`分支存在，则`default`分支代码块将得到执行； 否则，当前协程将被推入所有阻塞操作中相关的通道的发送数据协程队列或者接收数据协程队列中，并进入阻塞状态。

   

#### 源码分析

首先了解下channel是怎么创建的？

```go
func main() {
	c := make(chan int, 1)
	close(c)
}
```

通过`go tool compile -N -l -S main.go`输出其汇编代码，截取一小段观察一下：

```go
0x0024 00036 (main.go:9)        LEAQ    type.chan int(SB), AX // 将&chantype（元素类型是int）放到AX寄存器中
0x002b 00043 (main.go:9)        PCDATA  $2, $0
0x002b 00043 (main.go:9)        MOVQ    AX, (SP) // 也就是将&chantype放到SP（0）位置
0x002f 00047 (main.go:9)        MOVQ    $1, 8(SP)// 将1放到SP（8）位置
0x0038 00056 (main.go:9)        CALL    runtime.makechan(SB)// makechan(SP0,SP8)
0x003d 00061 (main.go:9)        PCDATA  $2, $1
0x003d 00061 (main.go:9)        MOVQ    16(SP), AX
0x0042 00066 (main.go:9)        MOVQ    AX, "".c+24(SP)
0x0047 00071 (main.go:10)       PCDATA  $2, $0
0x0047 00071 (main.go:10)       MOVQ    AX, (SP)
0x004b 00075 (main.go:10)       CALL    runtime.closechan(SB)
```

在上面的流程的关键部分加上了注释，也就是说咱们的`make(chan int, 1)`最终调用到了`runtime.makechan`这个方法。在进入分析之前先看看channel的结构：

```go
type hchan struct {
	qcount   uint           // 队列中实际有多少个元素
	dataqsiz uint           // channel的总长度（缓冲区的总长度）
	buf      unsafe.Pointer // 指向底层元素的指针
	elemsize uint16  // 元素类型的size
	closed   uint32  // 是否关闭，0：未关闭， 1：已关闭
	elemtype *_type // 元素类型
	sendx    uint   // 发送位置索引
	recvx    uint   // 接收位置索引
	recvq    waitq  // 接收者队列，一个双向链表
	sendq    waitq  // 发送者队列，一个双向链表

	lock mutex  // 锁，并发发送的时候需要上锁
}
```

咱们可以将`hchan`中的`buf`简单的看成一个数组缓冲区，`qcount`是数组中实际存储元素的数量，`dataqsiz`是数组的容量，`elemtype`是数组元素的类型。`sendx`和`recvx`分别是发送索引位置和接收索引位置，每次操作都会自增1，当`sendx`和`recv`等于`dataqsiz`时，会重置为零。`recvq`和`sendq`都是双向链表，里面维护着等待接收和等待发送的`goroutine`。当多个`gouroutine`并发操作同一个`channel`时，会使用`lock`进行控制。

##### 创建流程

```go
func makechan(t *chantype, size int) *hchan {
	elem := t.elem

	// 缓冲区中元素类型的尺寸不能超过16k
	if elem.size >= 1<<16 {
		throw("makechan: invalid channel element type")
	}
    // 判断是否位数对齐，
	if hchanSize%maxAlign != 0 || elem.align > maxAlign {
		throw("makechan: bad alignment")
	}
	// 计算缓冲区的总长度，并判断是否溢出
	mem, overflow := math.MulUintptr(elem.size, uintptr(size))
	if overflow || mem > maxAlloc-hchanSize || size < 0 {
		panic(plainError("makechan: size out of range"))
	}
	var c *hchan
	switch {
	case mem == 0:
		// channel长度或者元素类型尺寸为0时，也就是缓冲区长度为0时，只用分配hchan所占用的内存空间。
		c = (*hchan)(mallocgc(hchanSize, nil, true))
		// Race detector uses this location for synchronization.
		c.buf = c.raceaddr()
	case elem.kind&kindNoPointers != 0:
		// 元素类型不是指针类型，则将hchan和buf一次性分配出来
		c = (*hchan)(mallocgc(hchanSize+mem, nil, true))
        // 缓冲区buf的指针位置在c+hchanSize（hchanSize补齐为8的倍数）
		c.buf = add(unsafe.Pointer(c), hchanSize,hchanSize补齐为8的倍数)
	default:
		// 元素类型是指针类型，hchan和缓冲区单独分配
		c = new(hchan)
		c.buf = mallocgc(mem, elem, true)
	}
	// 元素的尺寸
	c.elemsize = uint16(elem.size)
	c.elemtype = elem
	c.dataqsiz = uint(size)


	return c
}
```

创建channel时共分为三种情况：

1. 缓冲区大小为0的情况下，只用给`hchan`分配内存即可。
2. 当元素类型不为指针时，可以考虑分配一段连续的内存，这样方便垃圾回收。
3. 当元素类型为指针时，需要给`hchan`和`buf`分别开辟空间。

最终都调用到`mallocgc`方法进行内存的分配，分配过程这里不做过多的描述，后面会考虑写一篇文章介绍一下分配的相关流程。



##### 发送流程

