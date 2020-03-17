package main

import "fmt"

/*
@Time : 2020/3/17 23:53
@Author : hejun
*/

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4
*/

func main() {
	fmt.Println(singleNumber1([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
loop:
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[j] == nums[i] {
				continue loop
			}
		}
		return nums[i]
	}
	return 0
}

func singleNumber1(nums []int) int {
	// 相同的数字异或之后都为0
	// 任何数字与0异或都是原本数字
	ret := 0
	for _, value := range nums {
		ret ^= value
	}
	return ret
}
