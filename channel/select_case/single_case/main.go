package main

/*
@Time : 2020/1/9
@Author : hejun
*/
func main() {
	c := make(chan int)
	close(c)
	select {
	case <-c:
	default:
	}
}

/**
0x0050 00080 (main.go:11)       MOVQ    $0, (SP)
0x0058 00088 (main.go:11)       PCDATA  $2, $1
0x0058 00088 (main.go:11)       PCDATA  $0, $0
0x0058 00088 (main.go:11)       MOVQ    "".c+24(SP), AX
0x005d 00093 (main.go:11)       PCDATA  $2, $0
0x005d 00093 (main.go:11)       MOVQ    AX, 8(SP)
0x0062 00098 (main.go:11)       CALL    runtime.selectnbrecv(SB)

*/
