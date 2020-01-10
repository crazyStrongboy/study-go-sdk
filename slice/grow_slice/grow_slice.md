slice的扩容过程：
```go
func main() {
	s := make([]int, 0, 1)
	s = append(s, 5,6)
}
```
slice的结构已经在simple_slice中有过介绍。

执行go tool compile -S -l -N main.go看一下汇编代码（移除了部分检测指令）：
```go
0x0000 00000 (main.go:8)        TEXT    "".main(SB), ABIInternal, $104-0
0x001a 00026 (main.go:8)        SUBQ    $104, SP
0x001e 00030 (main.go:8)        MOVQ    BP, 96(SP)
0x0023 00035 (main.go:8)        LEAQ    96(SP), BP
// make([]int, 0, 1)
0x0028 00040 (main.go:9)        MOVQ    $0, ""..autotmp_1+64(SP)
0x0031 00049 (main.go:9)        LEAQ    ""..autotmp_1+64(SP), AX
0x003c 00060 (main.go:9)        MOVQ    AX, "".s+72(SP)
0x0041 00065 (main.go:9)        MOVQ    $0, "".s+80(SP)
0x004a 00074 (main.go:9)        MOVQ    $1, "".s+88(SP)
// func growslice(et *_type, old slice, cap int) slice
0x0055 00085 (main.go:10)       LEAQ    type.int(SB), CX
0x005c 00092 (main.go:10)       MOVQ    CX, (SP) // 元素类型为int
0x0060 00096 (main.go:10)       MOVQ    AX, 8(SP) // 老的slice
0x0065 00101 (main.go:10)       MOVQ    $0, 16(SP)
0x006e 00110 (main.go:10)       MOVQ    $1, 24(SP)
0x0077 00119 (main.go:10)       MOVQ    $2, 32(SP)// 新的容量
0x0080 00128 (main.go:10)       CALL    runtime.growslice(SB)
// AX ===> newslice.array 
0x0085 00133 (main.go:10)       MOVQ    40(SP), AX
// CX ===> newslice.len
0x008a 00138 (main.go:10)       MOVQ    48(SP), CX
// DX ===> newslice.cap
0x008f 00143 (main.go:10)       MOVQ    56(SP), DX
0x0094 00148 (main.go:10)       ADDQ    $2, CX //CX = 2
0x009a 00154 (main.go:10)       MOVQ    $5, (AX) //newslice.array[0]= 5
0x00a1 00161 (main.go:10)       MOVQ    $6, 8(AX) //newslice.array[1] = 6
0x00a9 00169 (main.go:10)       MOVQ    AX, "".s+72(SP) //&autotmp_1
0x00ae 00174 (main.go:10)       MOVQ    CX, "".s+80(SP) //len=2
0x00b3 00179 (main.go:10)       MOVQ    DX, "".s+88(SP) //扩容后从cap
0x00b8 00184 (main.go:11)       MOVQ    96(SP), BP
0x00bd 00189 (main.go:11)       ADDQ    $104, SP
0x00c1 00193 (main.go:11)       RET
```

第一步操作 s:=make([]int, 0, 1)，在`append`的过程中，需要添加的元素的长度为2，但是当前slice容量为1，则需要进行扩容，咱们重点看看`runtime.growslice`方法。根据方法签名可以知道：第一个参数是元素类型，第二个参数是老的slice，第三个参数是新的容量。根据当前上下文可以知道当前元素类型是int，slice为变量s，新的容量为2。下面咱们去分析`growslice`的源码：

```go
func growslice(et *_type, old slice, cap int) slice {
    // 新分配的容量不能低于老的容量，否则，panic
	if cap < old.cap {
		panic(errorString("growslice: cap out of range"))
	}

	if et.size == 0 {
		// 对于元素尺寸为0的进行特殊处理，不用考虑拷贝老的元素到新slice中来
		return slice{unsafe.Pointer(&zerobase), old.len, cap}
	}
	// 计算新的slice的容量
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		if old.len < 1024 {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < cap {
				newcap += newcap / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = cap
			}
		}
	}

	var overflow bool // 用以判断是否会溢出
	var lenmem, newlenmem, capmem uintptr
    // 根据不同的元素尺寸计算出相应的指标，capmem需要进行位对齐
    // 1. 元素尺寸为1的，计算长度和容量不需要特殊处理
    // 2. 元素尺寸为8，即一个指针的长度时，用一个固定的偏移量去计算
    // 3. 元素尺寸为2的n次方，则进行位运算
    // 4. 其他情况下进行最简单的乘除运算即可
	switch {
	case et.size == 1:
		lenmem = uintptr(old.len)
		newlenmem = uintptr(cap)
		capmem = roundupsize(uintptr(newcap))
		overflow = uintptr(newcap) > maxAlloc
		newcap = int(capmem)
	case et.size == sys.PtrSize: 
		lenmem = uintptr(old.len) * sys.PtrSize
		newlenmem = uintptr(cap) * sys.PtrSize
		capmem = roundupsize(uintptr(newcap) * sys.PtrSize)
		overflow = uintptr(newcap) > maxAlloc/sys.PtrSize
		newcap = int(capmem / sys.PtrSize)
	case isPowerOfTwo(et.size):
		var shift uintptr
		if sys.PtrSize == 8 {
			// Mask shift for better code generation.
			shift = uintptr(sys.Ctz64(uint64(et.size))) & 63
		} else {
			shift = uintptr(sys.Ctz32(uint32(et.size))) & 31
		}
		lenmem = uintptr(old.len) << shift
		newlenmem = uintptr(cap) << shift
		capmem = roundupsize(uintptr(newcap) << shift)
		overflow = uintptr(newcap) > (maxAlloc >> shift)
		newcap = int(capmem >> shift)
	default:
		lenmem = uintptr(old.len) * et.size
		newlenmem = uintptr(cap) * et.size
		capmem, overflow = math.MulUintptr(et.size, uintptr(newcap))
		capmem = roundupsize(capmem)
		newcap = int(capmem / et.size)
	}
	// 判断是否溢出
	if overflow || capmem > maxAlloc {
		panic(errorString("growslice: cap out of range"))
	}

	var p unsafe.Pointer
	if et.kind&kindNoPointers != 0 {
        // 元素是非指针类型，直接分配noscan区域的span
		p = mallocgc(capmem, nil, false)
        // 进行缓存行填充后，会多出部分unreachable的内存，将其清理掉
		memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
	} else {
		p = mallocgc(capmem, et, true)
		if writeBarrier.enabled {
			bulkBarrierPreWriteSrcOnly(uintptr(p), uintptr(old.array), lenmem)
		}
	}
    // 拷贝old_slice中的元素到新的slice中
	memmove(p, old.array, lenmem)

	return slice{p, old.len, newcap}
}
```

源码不长，逻辑也很清晰，首先要计算出新`slice`的容量cap，然后计算出要分配的内存的大小并进行内存的分配，最后拷贝老的`slice`中的元素到新的`slice`中。

