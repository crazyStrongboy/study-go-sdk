package main

import "fmt"

/*
@Time : 2020/3/11
@Author : hejun
*/
/**
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 3
输出: [1,3,3,1]
*/
func main() {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	tmp := make([]int, rowIndex+1, rowIndex+1)
	z := 0
	for z <= rowIndex {
		tmp[z] = 1
		z++
	}
	for i := 2; i <= rowIndex; i++ {
		for j := i - 1; j > 0; j-- {
			tmp[j] = tmp[j] + tmp[j-1]
		}
	}
	return tmp
}
