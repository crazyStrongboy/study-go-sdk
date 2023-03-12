package main

import "fmt"

func maxCoins(nums []int) int {
	n := len(nums)
	// 先给数组前后都追加一个1
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	var dp [][]int
	for i := 0; i < n+2; i++ {
		dp = append(dp, make([]int, n+2))
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[k][j]+nums[i]*nums[k]*nums[j]+dp[i][k])
				fmt.Println(dp)
			}
		}
	}
	return dp[0][n+1]
}
