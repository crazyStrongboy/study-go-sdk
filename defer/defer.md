### 前言

`defer`这个关键字在开发过程中上场率可不低，初学者只会知道在当前函数中声明一个`defer`函数，那么会在当前函数`return`时再去执行`defer`定义的函数，但具体原因是什么呢？假如在当前函数中同时声明多个`defer`函数，为何先声明的后执行呢？要想弄懂这些问题，可以从本文中得到答案。

### Example

简单的看一段程序，代码很简单，咱们主要来分析一下`defer`具体的执行流程。一个简单的`defer`关键字在经过编译之后究竟经过了哪些方法的调用。

```go
func main() {
	f()
}

func f() {
	defer sum(1, 2)
}

func sum(a, b int) int {
	return a + b
}
```

咱们看一下汇编指令，重点部分都加有注释。（对函数调用过程这一块不太了解的可以去看一下[这篇文章](https://cloud.tencent.com/developer/article/1450282)）

```assembly
(gdb) disass main.f
Dump of assembler code for function main.f:                                
   0x000000000044f960 <0>:      mov    %fs:0xfffffffffffffff8,%rcx     
   0x000000000044f969 <+9>:     cmp    0x10(%rcx),%rsp                             
   0x000000000044f96d <+13>:    jbe    0x44f9cb <main.f+107>                       
   0x000000000044f96f <+15>:    sub    $0x30,%rsp # 分配函数f的栈空间
   0x000000000044f973 <+19>:    mov    %rbp,0x28(%rsp) # 保存main函数的rbp到0x28(%rsp)中
   0x000000000044f978 <+24>:    lea    0x28(%rsp),%rbp # 将0x28(%rsp)位置作为函数f的rbp
   0x000000000044f97d <+29>:    movl   $0x18,(%rsp)  # rsp(0)位置为24，代表参数长度     
   0x000000000044f984 <+36>:    lea    0x23c85(%rip),%rax   # 0x473610 这里是sum函数的地址   
   0x000000000044f98b <+43>:    mov    %rax,0x8(%rsp)  # rsp(8)位置为sum函数的地址     
   0x000000000044f990 <+48>:    movq   $0x1,0x10(%rsp)  # rsp(16)位置为第一个参数值1
   0x000000000044f999 <+57>:    movq   $0x2,0x18(%rsp) # rsp(24)位置为第二个参数值2
   0x000000000044f9a2 <+66>:    callq  0x4221f0 <runtime.deferproc># 关注点1 
   # deferproc后自动插入的一条指令。正常情况下eax寄存器中返回的值是0，执行接下来的业务逻辑。
   # 但是当业务逻辑panic后，并且有recover的情况下，eax寄存器中会被填入1，
   # 则经过该指令对比后直接跳转到0x44f9bb处的deferreturn
   0x000000000044f9a7 <+71>:    test   %eax,%eax  
   0x000000000044f9a9 <+73>:    jne    0x44f9bb <main.f+91>                         
   0x000000000044f9ab <+75>:    nop                                                 
   0x000000000044f9ac <+76>:    callq  0x422a80 <runtime.deferreturn> # 关注点2
   0x000000000044f9b1 <+81>:    mov    0x28(%rsp),%rbp # 恢复成main函数的rbp       
   0x000000000044f9b6 <+86>:    add    $0x30,%rsp   # 将栈空间还原
   0x000000000044f9ba <+90>:    retq    
   0x000000000044f9bb <+91>:	nop
   0x000000000044f9bc <+92>:	callq  0x422a80 <runtime.deferreturn>
   0x000000000044f9c1 <+97>:	mov    0x28(%rsp),%rbp
   0x000000000044f9c6 <+102>:	add    $0x30,%rsp
   0x000000000044f9ca <+106>:	retq   
   0x000000000044f9cb <+107>:	callq  0x447840 <runtime.morestack_noctxt>
   0x000000000044f9d0 <+112>:	jmp    0x44f960 <main.f>

```



上面有两个关注点，`runtime.deferproc`与`runtime.deferreturn`这两个方法。在了解这两个方法之前，先看一下`defer`的结构。

```go
// A _defer holds an entry on the list of deferred calls.
// If you add a field here, add code to clear it in freedefer.
type _defer struct {
	siz     int32 // 参数的长度，函数fn的参数长度
	started bool
	sp      uintptr // sp at time of defer
	pc      uintptr // defer语句下一条语句的地址
    fn      *funcval // 按上面的例子是&funcval{fn:&sum}
	_panic  *_panic // panic that is running defer
	link    *_defer // 同一个goroutine所有被延迟执行的函数通过该成员链在一起形成一个链表
}
```

### deferproc

```go
// siz=fn.fn的参数长度
func deferproc(siz int32, fn *funcval) { // arguments of fn follow fn
	if getg().m.curg != getg() {
		// go code on the system stack can't defer
		throw("defer on system stack")
	}

	// 获取调用者的sp值（栈顶值）
	sp := getcallersp()
	argp := uintptr(unsafe.Pointer(&fn)) + unsafe.Sizeof(fn)
    // 获取调用者的pc
	callerpc := getcallerpc()

    // 从缓冲区获取或者重新创建一个
	d := newdefer(siz)
	if d._panic != nil {
		throw("deferproc: d.panic != nil after newdefer")
	}
	d.fn = fn
	d.pc = callerpc // 在panic-recover后，用来跳转的指令
	d.sp = sp
	switch siz {
	case 0:
		// 表示参数长度是0，无需进行任何操作
	case sys.PtrSize:
        // 表示是指针参数，直接拷贝地址值即可
		*(*uintptr)(deferArgs(d)) = *(*uintptr)(unsafe.Pointer(argp))
	default:
        // 相当于复制值到defer之后
		memmove(deferArgs(d), unsafe.Pointer(argp), uintptr(siz))
	}

	// 正常情况下是返回0，然后执行defer后面的逻辑，最后在f中执行return时调用deferreturn
    // 异常情况下（panic-recover）返回1，直接执行deferreturn
	return0()
}
```

首先，`deferproc`需要两个参数，第一个是defer函数的参数的长度（以字节为单位的），第二个参数 funcval 是一个变长结构体。如下所示：

```go
type funcval struct {
     fn uintptr 
    // variable-size, fn-specific data here
}
```

`fn`的参数部分紧随着`fn`，按本文的例子，这里的`fn=&sum`，参数a,b以及返回参数紧跟着`fn`排列。由于go中函数调用参数通过栈来传递，所以此时堆栈结构是：

![deferproc](http://images.hcyhj.cn/blogimages/defer/deferproc.png)

在函数`deferproc`中，

1. 会先通过`newdefer(siz)`获取一个`defer`结构体对应的对象。具体是从缓存中获取还是新分配一个，下面会进行详细的解释。
2. 并给其sp属性附上当前调用者f的栈顶sp的值，后面进行`deferreturn`时会通过这个值去进行判断要执行的`defer`是否属于当前调用者。
3. 然后会将参数部分拷贝到紧挨着`defer`对象后面的地址：`deferArgs(d)=unsafe.Pointer(d)+unsafe.Sizeof(*d)`。
4. 执行`return0`函数，正常情况下返回0，经过`test   %eax,%eax`检测后继续执行业务逻辑。异常情况下会返回1，并且直接跳转到`deferreturn`。



### newdefer

```go
//go:nosplit
func newdefer(siz int32) *_defer {
	var d *_defer
    // 计算出sc，方便从p的缓存池或者sched全局缓冲池中获取defer
	sc := deferclass(uintptr(siz))
	gp := getg()
	if sc < uintptr(len(p{}.deferpool)) {
		pp := gp.m.p.ptr()
		if len(pp.deferpool[sc]) == 0 && sched.deferpool[sc] != nil {
			// 当前p上缓存用完了，则需要从全局区拷贝几个出来，
            // 直到拷贝到deferpool容量的一半
			systemstack(func() {
                // 加上锁，有可能多个p同时去拉取
				lock(&sched.deferlock)
				for len(pp.deferpool[sc]) < cap(pp.deferpool[sc])/2 && sched.deferpool[sc] != nil {
					d := sched.deferpool[sc]
					sched.deferpool[sc] = d.link
					d.link = nil
					pp.deferpool[sc] = append(pp.deferpool[sc], d)
				}
				unlock(&sched.deferlock)
			})
		}
		if n := len(pp.deferpool[sc]); n > 0 {
			d = pp.deferpool[sc][n-1]
			pp.deferpool[sc][n-1] = nil
			pp.deferpool[sc] = pp.deferpool[sc][:n-1]
		}
	}
	if d == nil {
		// Allocate new defer+args.
        // 当前p和全局缓冲池中都没有或者需要的参数过长，则需要新分配一个
		systemstack(func() {
			total := roundupsize(totaldefersize(uintptr(siz)))
			d = (*_defer)(mallocgc(total, deferType, true))
		})
		if debugCachedWork {
			// Duplicate the tail below so if there's a
			// crash in checkPut we can tell if d was just
			// allocated or came from the pool.
			d.siz = siz
			d.link = gp._defer
			gp._defer = d
			return d
		}
	}
	d.siz = siz
    // 与之前绑定在g上的defer形成一个链表
    // 例如之前g上绑定1号defer,新加进来一个2号defer
    // 那么现在g上defer的顺序是 2-->1,这也变向的验证了先声明的defer后执行
	d.link = gp._defer
	gp._defer = d
	return d
}
```

分配`defer`的流程大概分为三步走：

1. 由于缓冲区中是根据参数的长度进行不同级别的`defer`来缓存的，所以得先计算出相应的`sc`值。
2. 判断`sc`的值是否在缓冲区允许的范围内。
   - 在范围内，即sc<5，则先检查当前p的缓冲区中是否还有可用的`defer`。
     - 没有可用，则去全局缓冲区`sched`中批量拉取直至达到当前p缓冲区容量的一半。再从当前p中去获取。
     - 若有可用，则直接取出当前p缓冲区对应尺寸的`[]defer`最后一个元素即可。
3. 无法从缓冲区中获取或者是参数过大（sc不在缓冲区范围内），则直接新分配一个。
4. 将获取的`defer`绑定到当前goroutine上，并与之前绑定的`defer`形成链表。



### deferreturn

```go
func deferreturn(arg0 uintptr) {
	gp := getg()
    // 获取g上绑定的第一个defer
	d := gp._defer
	if d == nil {
        // 由于是递归调用，这里是一个循环终止条件，d上已经没有绑定的defer了
		return
	}
    // 获取当前调用者的sp
	sp := getcallersp()
	if d.sp != sp {
        // 判断当前调用者栈是否和defer中保存的一致
        // 举个例子，a()中声明一个defer1，并调用b(),b中也声明一个defer2
        // 然后defer1和defer2都绑定在同一个g上
        // 那么在b()执行return时，只会执行defer2，因为defer2上绑定的才是b()的sp
		return
	}
	// Moving arguments around.
	//
	// Everything called after this point must be recursively
	// nosplit because the garbage collector won't know the form
	// of the arguments until the jmpdefer can flip the PC over to
	// fn.
	switch d.siz {
	case 0:
		// Do nothing.
	case sys.PtrSize:
		*(*uintptr)(unsafe.Pointer(&arg0)) = *(*uintptr)(deferArgs(d))
	default:
		memmove(unsafe.Pointer(&arg0), deferArgs(d), uintptr(d.siz))
	}
	fn := d.fn
	d.fn = nil
    // g中的defer指向下一个defer
	gp._defer = d.link
    // 进行释放，归还到相应的缓冲区或者让gc回收
	freedefer(d)
	// 执行defer中绑定的func
	jmpdefer(fn, uintptr(unsafe.Pointer(&arg0)))
}
```

1. 判断当前goroutine上是否还有绑定的`defer`，若没有，直接return。
2. 获取goroutine绑定的`defer`链表头部的`defer`。
3. 判断当前`defer`中存储的sp是否和调用者的sp一致，若不一致，也直接return，证明当前defer不是在此调用函数中声明的。
4. 进行参数的拷贝。
5. 释放当前要执行fn关联的defer。
6. 执行`jmpdefer`函数，这里会执行完fn的逻辑后递归调用`deferreturn`函数。



接着看一下`jmpdefer`函数，咱们拆开一句句的去看。

*runtime/asm_amd64.s：566*

```assembly
// func jmpdefer(fv *funcval, argp uintptr)
// argp is a caller SP.
// called from deferreturn.
// 1. pop the caller
// 2. sub 5 bytes from the callers return
// 3. jmp to the argument
TEXT runtime·jmpdefer(SB), NOSPLIT, $0-16
	MOVQ	fv+0(FP), DX	// fn
	MOVQ	argp+8(FP), BX	// caller sp
	LEAQ	-8(BX), SP	// caller sp after CALL
	MOVQ	-8(SP), BP	// restore BP as if deferreturn returned (harmless if framepointers not in use)
	SUBQ	$5, (SP)	// return to CALL again
	MOVQ	0(DX), BX
	JMP	BX	// but first run the deferred function

```

![jmpdefer](http://images.hcyhj.cn/blogimages/defer/jmpdefer.png)

在执行`deferreturn`调用时，栈帧内部结构大概如上图所示，咱们简单的梳理一下。`deferreturn`只有一个参数`arg0`，所以在栈中的具体位置肯定是靠近其返回地址的。返回地址对应的sp实际上是函数f的sp。

这里咱们将上面的汇编代码拆解进行一一分析：

```assembly
MOVQ	fv+0(FP), DX   
```

将`fv`中的第一个参数`fn`放到DX寄存器中。

```assembly
MOVQ	argp+8(FP), BX
```

将`argp`的地址放到BX寄存器中，实际上就是上图中的`&arg0`。

```assembly
LEAQ	-8(BX), SP
```

将BX地址减去8，并将对应的地址放到寄存器SP中。实际上`&arg0`地址减去8指向的是`deferreturn`的返回地址，也就是` callq  0x422a80 <runtime.deferreturn>`的下一条指令`mov    0x28(%rsp),%rbp`，所以这里实际上将SP寄存器指向函数f的栈顶。

```assembly
MOVQ	-8(SP), BP
```

SP-8的位置恰好存放函数f的bp，这里将其值存放到BP寄存器中。经过了上面四句指令，已经成功将函数栈从`deferreturn`切换到了`f`中。

```assembly
SUBQ	$5, (SP)
```

将SP向上移动5位，一般人可能会在这里有疑惑，咱们先看一张图片

![jmpdefer2](http://images.hcyhj.cn/blogimages/defer/jmpdefer2.png)

咱们上面有说到，sp指向的位置是`deferreturn`的返回地址，也就是` callq  0x422a80 <runtime.deferreturn>`的下一条指令`mov    0x28(%rsp),%rbp`。很神奇的就是SP-5后，咱们可以发现SP又指向了`callq  0x422a80 <runtime.deferreturn>`的位置，相当于该位置成为了栈顶。这样在执行完函数`fn`之后，又会继续执行`deferreturn`函数，相当于一个递归调用。

```assembly
MOVQ	0(DX), BX #将fn地址放到BX寄存器中
JMP	BX
```

跳转到指定的fn函数去执行相关逻辑，执行完成后跳转到`deferreturn`函数。

### freedefer

```go
//go:nosplit
func freedefer(d *_defer) {
	if d._panic != nil {
		freedeferpanic()
	}
	if d.fn != nil {
		freedeferfn()
	}
	sc := deferclass(uintptr(d.siz))
	if sc >= uintptr(len(p{}.deferpool)) {
        // 参数过大的不进行缓存，等gc进行回收
		return
	}
	pp := getg().m.p.ptr()
	if len(pp.deferpool[sc]) == cap(pp.deferpool[sc]) {
		// 当前p中缓冲区已满，则迁移一半的defer到全局缓冲区
		systemstack(func() {
			var first, last *_defer
			for len(pp.deferpool[sc]) > cap(pp.deferpool[sc])/2 {
				n := len(pp.deferpool[sc])
				d := pp.deferpool[sc][n-1]
				pp.deferpool[sc][n-1] = nil
				pp.deferpool[sc] = pp.deferpool[sc][:n-1]
				if first == nil {
					first = d
				} else {
					last.link = d
				}
				last = d
			}
			lock(&sched.deferlock)
			last.link = sched.deferpool[sc]
			sched.deferpool[sc] = first
			unlock(&sched.deferlock)
		})
	}

	// These lines used to be simply `*d = _defer{}` but that
	// started causing a nosplit stack overflow via typedmemmove.
	d.siz = 0
	d.started = false
	d.sp = 0
	d.pc = 0
	// d._panic and d.fn must be nil already.
	// If not, we would have called freedeferpanic or freedeferfn above,
	// both of which throw.
	d.link = nil
	// 缓存到当前p的缓冲区
	pp.deferpool[sc] = append(pp.deferpool[sc], d)
}
```

释放已经用过的`defer`：

1. 首先判断参数长度，参数长度过长的直接让gc进行回收即可，无需归还到缓冲区中。
2. 若当前p的缓冲区已经满了，则需要进行迁移操作，这里会将当前p容量一半的`defer`归还到全局缓冲区，供给其他的p使用。操作的时候需要加上锁，防止多个p出现并发操作。
3. 将其属性置空，并追加到`pp.deferpool[sc] `数组中。



### 总结

本文主要涉及到`defer`这个关键字在编译之后究竟是怎样嵌入我们的代码的，以及`defer`后的函数是何时调用的，具体流程会先调用`deferproc`在goroutine上绑定`defer`链表，然后执行`deferreturn`时依次遍历链表执行`defer`中的函数，正向插入，反向遍历，这样也能看出先定义的`defer`后执行。当然还有部分点未做涉及，例如panic后当前goroutine`defer`链表中函数的调用，以及`recover`是如何实现的。下一篇文章会对这一块进行分析。



### 参考文章

[深入理解defer（下）defer实现机制](https://cloud.tencent.com/developer/article/1450260)