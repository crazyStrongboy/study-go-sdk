package main

import "fmt"

/*
@Time : 2020/3/7
@Author : hejun
*/

/**
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2

*/

func main() {
	fmt.Println(mySqrt1(9))
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	i := 2
	for {
		curr := x / i
		if curr < i {
			return i - 1
		}
		i++
	}
}

func mySqrt1(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	start := 0
	end := x/2 + 1
	for start < end {
		mid := (start + end + 1) >> 1
		if mid*mid > x {
			end = mid - 1
		} else {
			start = mid
		}
	}
	return start
}
