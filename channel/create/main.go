package main

/*
@Time : 2020/1/7
@Author : hejun
*/

func main() {
	c := make(chan int, 1)
	close(c)
}
