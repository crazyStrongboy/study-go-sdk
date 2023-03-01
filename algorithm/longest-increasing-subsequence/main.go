package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}

func lengthOfLIS(nums []int) int {
	result := 1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
		//fmt.Println(dp)
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
