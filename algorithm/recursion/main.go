package main

import "fmt"

/*
@Time : 2020/1/19
@Author : hejun
*/

var c map[int]int

func init() {
	c = map[int]int{}
}

func main() {
	fmt.Println(f(5))
}

func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if r, ok := c[n]; ok {
		return r
	}
	ret := f(n-1) + f(n-2)
	c[n] = ret // 避免重复计算
	return ret
}
