package main

import "fmt"

/*
@Time : 2020/1/19
@Author : hejun
*/

func main() {
	c := generateNaturalNumber()
	for i := 1; i <= 5; i++ {
		prime := <-c
		fmt.Printf("index: %d,prime: %d\n", i, prime)
		c = primeFilter(c, prime)
	}
}

// 从2开始生成自然数
func generateNaturalNumber() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 进行过滤
func primeFilter(oldC <-chan int, prime int) chan int {
	newC := make(chan int)
	go func() {
		for {
			if i := <-oldC; i%prime != 0 {
				newC <- i
			} else {
				fmt.Println(i)
			}

		}
	}()
	return newC
}
