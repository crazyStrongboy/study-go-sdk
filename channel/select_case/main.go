package main

/*
@Time : 2020/1/9
@Author : hejun
*/

func main() {
	c := make(chan int)
	a := make(chan int)
	close(c)
	close(a)
	select {
	case <-c:
	case <-a:
	default:

	}
}

/**
0x011d 00285 (main.go:13)       MOVQ    AX, (SP)
0x0121 00289 (main.go:13)       PCDATA  $2, $1
0x0121 00289 (main.go:13)       PCDATA  $0, $0
0x0121 00289 (main.go:13)       MOVQ    ""..autotmp_7+88(SP), AX
0x0126 00294 (main.go:13)       PCDATA  $2, $0
0x0126 00294 (main.go:13)       MOVQ    AX, 8(SP)
0x012b 00299 (main.go:13)       MOVQ    $3, 16(SP)
0x0134 00308 (main.go:13)       CALL    runtime.selectgo(SB)
*/
