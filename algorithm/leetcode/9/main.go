package main

import "fmt"

/*
@Time : 2020/3/5
@Author : hejun
*/
/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
*/

func main() {
	x := 0
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x%10 == 0 || x < 0 {
		return false
	}
	y := 0
	tmp := x
	for x != 0 {
		pop := x % 10
		y = y*10 + pop
		x = x / 10
	}
	if y == tmp {
		return true
	}
	return false
}
