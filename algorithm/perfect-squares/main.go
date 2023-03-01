package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(12))
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i < j*j {
				break
			}
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
		//fmt.Println(dp)
	}
	return dp[n]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
