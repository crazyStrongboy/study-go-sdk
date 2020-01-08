package main

/*
@Time : 2020/1/8
@Author : hejun
*/
func main() {
	c := make(chan int, 1)
	close(c)
	<-c
}

/**
  0x0055 00085 (main.go:10)       MOVQ    AX, (SP)
  0x0059 00089 (main.go:10)       MOVQ    $0, 8(SP)
  0x0062 00098 (main.go:10)       CALL    runtime.chanrecv1(SB)
  0x0067 00103 (main.go:11)       MOVQ    32(SP), BP
  0x006c 00108 (main.go:11)       ADDQ    $40, SP
*/
