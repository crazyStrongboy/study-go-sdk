package _6

/*
@Time : 2020/1/5 18:32
@Author : hejun
*/
func foo(c <-chan int) {
	close(c) // error: 不能关闭单向接收通道
}
