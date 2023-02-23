package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{{1, 2}, {3, 4}}))
}

func minPathSum(grid [][]int) int {
	var dp [][]int
	m := len(grid)
	n := len(grid[0])
	dp = append(dp, make([]int, n))
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp = append(dp, make([]int, n))
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
