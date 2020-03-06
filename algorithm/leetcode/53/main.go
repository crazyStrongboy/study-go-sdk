package main

import "fmt"

/*
@Time : 2020/3/6
@Author : hejun
*/

/**

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

func main() {
	max := maxSubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(max)
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		count := nums[i]
		if count > max {
			max = count
		}
		for j := i + 1; j < len(nums); j++ {
			count += nums[j]
			if count > max {
				max = count
			}
		}
	}
	return max
}

// 动态规划
func maxSubArray1(nums []int) int {
	n := len(nums)
	maxNum := nums[0]
	for i := 1; i < n; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		maxNum = max(nums[i], maxNum)
	}
	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
