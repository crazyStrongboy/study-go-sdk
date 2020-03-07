package main

import "fmt"

/*
@Time : 2020/3/7
@Author : hejun
*/

/**

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
*/

func main() {
	fmt.Println(climbStairs2(44))
}

// 斐波那契数
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

// 记忆化递归
func climbStairs2(n int) int {
	arr := make([]int, n+1, n+1)
	i := doClimb(0, n, arr)
	fmt.Println(arr)
	return i
}

func doClimb(i int, n int, arr []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if arr[i] > 0 {
		return arr[i]
	}
	arr[i] = doClimb(i+1, n, arr) + doClimb(i+2, n, arr)
	return arr[i]
}

// 递归的算法
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	// 第一次爬一阶
	// 第一次爬二阶
	return climbStairs(n-1) + climbStairs(n-2)

}
