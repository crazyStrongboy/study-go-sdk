package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2}))
}

// dp[i][j] 含义：  第i天 j类型操作
// 状态表
// 0 买入
// 1 卖出
// 2 冷冻期
func maxProfit(prices []int) int {
	var dp [][]int
	for i := 0; i < len(prices); i++ {
		dp = append(dp, make([]int, 3))
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][2]-prices[i], dp[i-1][0])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = dp[i-1][1]
		for _, ints := range dp {
			fmt.Println(ints)
		}
		fmt.Println()
	}
	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
