package main

import "fmt"

/*
@Time : 2020/3/11
@Author : hejun
*/

/**
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

func main() {
	ints := generate(5)
	fmt.Println(ints)
}

func generate(numRows int) [][]int {
	result := make([][]int, 0, numRows)
	if numRows == 0 {
		return result
	}
	first := []int{1}
	result = append(result, first)
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1, i+1)
		pre := result[i-1]
		row[0] = 1
		for j := 1; j < i; j++ {
			row[j] = pre[j-1] + pre[j]
		}
		row[len(row)-1] = 1
		result = append(result, row)
	}
	return result
}
