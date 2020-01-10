看一下简单的slice，不进行扩容

```go
func main() {
	s := make([]int, 0, 1)
	s = append(s, 5)
}
```

首先看看slice的数据结构：

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```



执行go tool compile -S -l -N main.go看一下汇编代码（移除了部分检测指令）：

```go
0x0000 00000 (main.go:7)        TEXT    "".main(SB), NOSPLIT|ABIInternal, $40-0
0x0000 00000 (main.go:7)        SUBQ    $40, SP 
0x0004 00004 (main.go:7)        MOVQ    BP, 32(SP)
0x0009 00009 (main.go:7)        LEAQ    32(SP), BP 
0x000e 00014 (main.go:8)        MOVQ    $0, ""..autotmp_1(SP)
0x0016 00022 (main.go:8)        LEAQ    ""..autotmp_1(SP), AX
0x0020 00032 (main.go:8)        MOVQ    AX, "".s+8(SP)
0x0025 00037 (main.go:8)        MOVQ    $0, "".s+16(SP)
0x002e 00046 (main.go:8)        MOVQ    $1, "".s+24(SP)
0x0039 00057 (main.go:9)        MOVQ    $5, ""..autotmp_1(SP)
0x0041 00065 (main.go:9)        MOVQ    AX, "".s+8(SP)
0x0046 00070 (main.go:9)        MOVQ    $1, "".s+16(SP)
0x004f 00079 (main.go:9)        MOVQ    $1, "".s+24(SP)
0x0058 00088 (main.go:10)       MOVQ    32(SP), BP
0x005d 00093 (main.go:10)       ADDQ    $40, SP
0x0061 00097 (main.go:10)       RET
```

咱们简单的组合一下上面的指令：

```go
SUBQ    $40, SP
	MOVQ    BP, 32(SP)
	LEAQ    32(SP), BP
		MOVQ    $0, ""..autotmp_1(SP) //autotmp_1 = 0; autotmp_1是一个临时变量
		LEAQ    ""..autotmp_1(SP), AX //AX = &autotmp_1
		MOVQ    AX, "".s+8(SP) // array = &autotmp_1
		MOVQ    $0, "".s+16(SP) // len = 0
		MOVQ    $1, "".s+24(SP)  //cap = 1
		// 上面5条指令完成了make([]int, 0, 1)的创建。

		MOVQ    $5, ""..autotmp_1(SP)// autotmp_1[0] = 5
		MOVQ    AX, "".s+8(SP) // &autotmp_1
		MOVQ    $1, "".s+16(SP) //len = 1
		MOVQ    $1, "".s+24(SP) // cap = 1
	MOVQ    32(SP), BP
ADDQ    $40, SP
```

共分为三层：

第一层的`SUBQ    $40, SP`与`ADDQ    $40, SP`，这两条指令分别是分配栈针的长度以及回收栈针。

第二层是保存BP现场，多分配了8个长度的栈针去保存BP的地址。

第三层是具体执行`main`函数中的代码块，具体执行情况看上面的注释。可参考下面两个示意图：

make([]int, 0, 1)流程：

![make_slice](http://images.hcyhj.cn/blogimages/slice/make_slice.png)



append(s, 5)流程：

![make_slice](http://images.hcyhj.cn/blogimages/slice/simple_append.png))

