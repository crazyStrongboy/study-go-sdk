package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	i := 0
	for i <= amount {
		dp[i] = math.MaxInt
		i++
	}
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt {
				dp[j] = min(dp[j-coins[i]]+1, dp[j])
			}
			fmt.Println(dp)
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
