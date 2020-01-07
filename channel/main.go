package main

import "fmt"

/*
@Time : 2020/1/7
@Author : hejun
*/

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	t := <-c
	fmt.Println(t)
	close(c)
	//selectCase()
	forRange()
}

func forRange() {
	c := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		c <- i
	}
	close(c)
	for v := range c {
		fmt.Println(v)
	}
}

func selectCase() {
	c := make(chan int, 1)
	c <- 1
	close(c)
	select {
	case <-c:
		fmt.Println("xxxx")
	default:
		fmt.Println("aaaa")
	}
}
