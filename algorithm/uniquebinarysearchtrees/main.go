package main

import "fmt"

func main() {

}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	//dp[1] =1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	fmt.Println(dp)
	return dp[n]
}
