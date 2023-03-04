package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{14, 9, 8, 4, 3, 2}))
}

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
	if sum%2 != 0 {
		return false
	}
	return bag(nums, sum/2)
}

func bag(nums []int, weight int) bool {
	dp := make([]int, weight+1)
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := weight; j > 0; j-- {
			if j >= nums[i] {
				dp[j] = max(dp[j-nums[i]]+nums[i], dp[j])
			}
		}
		//fmt.Println(dp)
	}
	return dp[weight] == weight
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func backtrack(nums []int, start, target int) bool {
	if target < 0 {
		return false
	}
	if target == 0 {
		return true
	}
	for i := start; i < len(nums); i++ {
		if backtrack(nums, i+1, target-nums[i]) {
			return true
		}
	}
	return false
}
