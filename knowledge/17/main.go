package main

/*
@Time : 2020/1/5 18:37
@Author : hejun
*/

func main() {
	c := make(chan int)
	close(c)
	c <- 1 //panic: send on closed channel
}
