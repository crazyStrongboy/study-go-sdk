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
	}
}

/**
0x005a 00090 (main.go:11)       MOVQ    "".c+24(SP), AX
0x005f 00095 (main.go:11)       PCDATA  $2, $0
0x005f 00095 (main.go:11)       MOVQ    AX, (SP)
0x0063 00099 (main.go:11)       MOVQ    $0, 8(SP)
0x006c 00108 (main.go:11)       CALL    runtime.chanrecv1(SB)
*/
