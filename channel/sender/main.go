package main

/*
@Time : 2020/1/8
@Author : hejun
*/
func main() {
	c := make(chan int, 1)
	c <- 1
	close(c)
}

/**
0x0057 00087 (main.go:9)        CALL    runtime.chansend1(SB)
0x005c 00092 (main.go:10)       PCDATA  $2, $1
0x005c 00092 (main.go:10)       PCDATA  $0, $0
0x005c 00092 (main.go:10)       MOVQ    "".c+24(SP), AX
0x0061 00097 (main.go:10)       PCDATA  $2, $0
0x0061 00097 (main.go:10)       MOVQ    AX, (SP)
0x0065 00101 (main.go:10)       CALL    runtime.closechan(SB)
*/
