package main

/*
@Time : 2020/1/9
@Author : hejun
*/

func main() {
	c := make(chan int)
	close(c)
}

/**
  0x0038 00056 (main.go:9)        CALL    runtime.makechan(SB)
  0x003d 00061 (main.go:9)        PCDATA  $2, $1
  0x003d 00061 (main.go:9)        MOVQ    16(SP), AX
  0x0042 00066 (main.go:9)        MOVQ    AX, "".c+24(SP)
  0x0047 00071 (main.go:10)       PCDATA  $2, $0
  0x0047 00071 (main.go:10)       MOVQ    AX, (SP)
  0x004b 00075 (main.go:10)       CALL    runtime.closechan(SB)
*/
