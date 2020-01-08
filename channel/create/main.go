package main

/*
@Time : 2020/1/7
@Author : hejun
*/

func main() {
	c := make(chan int, 1)
	close(c)
}

/**
go tool compile -N -l -S main.go

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
*/
