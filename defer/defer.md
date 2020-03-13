



简单的看一段程序，代码很简单，咱们主要来分析一下`defer`具体的执行流程。

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

咱们看一下汇编指令，重点部分都加有注释，对函数调用过程这一块不太了解的可以去看一下[这篇文章](https://cloud.tencent.com/developer/article/1450282)。

```assembly
x0x44f94b <main.main+27> callq  0x44f960 <main.f>                                         
x0x44f950 <main.main+32> mov    (%rsp),%rbp                                               
x0x44f954 <main.main+36> add    $0x8,%rsp                                                 
x0x44f958 <main.main+40> retq                                                             
x0x44f959 <main.main+41> callq  0x447840 <runtime.morestack_noctxt>                       
x0x44f95e <main.main+46> jmp    0x44f930 <main.main>                                     
x0x44f960 <main.f>       mov    %fs:0xfffffffffffffff8,%rcx                               
x0x44f969 <main.f+9>     cmp    0x10(%rcx),%rsp                                           
x0x44f96d <main.f+13>    jbe    0x44f9cb <main.f+107>                                     
x0x44f96f <main.f+15>    sub    $0x30,%rsp # 分配函数f的栈空间
x0x44f973 <main.f+19>    mov    %rbp,0x28(%rsp) # 保存main函数的rbp到0x28(%rsp)中
x0x44f978 <main.f+24>    lea    0x28(%rsp),%rbp # 将0x28(%rsp)位置作为函数f的rbp
x0x44f97d <main.f+29>    movl   $0x18,(%rsp)  # rsp(0)位置为24，代表参数长度           
x0x44f984 <main.f+36>    lea    0x23c85(%rip),%rax   # 0x473610 这里是sum函数的地址    
x0x44f98b <main.f+43>    mov    %rax,0x8(%rsp)  # rsp(8)位置为sum函数的地址      
x0x44f990 <main.f+48>    movq   $0x1,0x10(%rsp)  # rsp(16)位置为第一个参数值1               
x0x44f999 <main.f+57>    movq   $0x2,0x18(%rsp) # rsp(24)位置为第一个参数值2
x0x44f9a2 <main.f+66>    callq  0x4221f0 <runtime.deferproc># 关注点1 
x0x44f9a7 <main.f+71>    test   %eax,%eax                                                 
x0x44f9a9 <main.f+73>    jne    0x44f9bb <main.f+91>                                     
x0x44f9ab <main.f+75>    nop                                                             
x0x44f9ac <main.f+76>    callq  0x422a80 <runtime.deferreturn> # 关注点2
x0x44f9b1 <main.f+81>    mov    0x28(%rsp),%rbp # 恢复成main函数的rbp       
x0x44f9b6 <main.f+86>    add    $0x30,%rsp   # 将栈空间还原
x0x44f9ba <main.f+90>    retq     
```



关注一下上面的两个关注点，`runtime.deferproc`与`runtime.deferreturn`这两个方法。



```go
// A _defer holds an entry on the list of deferred calls.
// If you add a field here, add code to clear it in freedefer.
type _defer struct {
	siz     int32 // 参数的长度
	started bool
	sp      uintptr // sp at time of defer
	pc      uintptr
	fn      *funcval
	_panic  *_panic // panic that is running defer
	link    *_defer
}
```









```go
func deferproc(siz int32, fn *funcval) { // arguments of fn follow fn
	if getg().m.curg != getg() {
		// go code on the system stack can't defer
		throw("defer on system stack")
	}

	// the arguments of fn are in a perilous state. The stack map
	// for deferproc does not describe them. So we can't let garbage
	// collection or stack copying trigger until we've copied them out
	// to somewhere safe. The memmove below does that.
	// Until the copy completes, we can only call nosplit routines.
	sp := getcallersp()
	argp := uintptr(unsafe.Pointer(&fn)) + unsafe.Sizeof(fn)
	callerpc := getcallerpc()

	d := newdefer(siz)
	if d._panic != nil {
		throw("deferproc: d.panic != nil after newdefer")
	}
	d.fn = fn
	d.pc = callerpc
	d.sp = sp
	switch siz {
	case 0:
		// Do nothing.
	case sys.PtrSize:
		*(*uintptr)(deferArgs(d)) = *(*uintptr)(unsafe.Pointer(argp))
	default:
		memmove(deferArgs(d), unsafe.Pointer(argp), uintptr(siz))
	}

	// deferproc returns 0 normally.
	// a deferred func that stops a panic
	// makes the deferproc return 1.
	// the code the compiler generates always
	// checks the return value and jumps to the
	// end of the function if deferproc returns != 0.
	return0()
	// No code can go here - the C return register has
	// been set and must not be clobbered.
}
```





```go
//go:nosplit
func newdefer(siz int32) *_defer {
	var d *_defer
	sc := deferclass(uintptr(siz))
	gp := getg()
	if sc < uintptr(len(p{}.deferpool)) {
		pp := gp.m.p.ptr()
		if len(pp.deferpool[sc]) == 0 && sched.deferpool[sc] != nil {
			// Take the slow path on the system stack so
			// we don't grow newdefer's stack.
			systemstack(func() {
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
	d.link = gp._defer
	gp._defer = d
	return d
}
```







```go
func deferreturn(arg0 uintptr) {
	gp := getg()
	d := gp._defer
	if d == nil {
		return
	}
	sp := getcallersp()
	if d.sp != sp {
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
	gp._defer = d.link
	freedefer(d)
	jmpdefer(fn, uintptr(unsafe.Pointer(&arg0)))
}
```



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
		return
	}
	pp := getg().m.p.ptr()
	if len(pp.deferpool[sc]) == cap(pp.deferpool[sc]) {
		// Transfer half of local cache to the central cache.
		//
		// Take this slow path on the system stack so
		// we don't grow freedefer's stack.
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

	pp.deferpool[sc] = append(pp.deferpool[sc], d)
}
```



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
	MOVQ	-8(SP), BP	// 
	SUBQ	$5, (SP)	// return to CALL again
	MOVQ	0(DX), BX
	JMP	BX	// but first run the deferred function

```







