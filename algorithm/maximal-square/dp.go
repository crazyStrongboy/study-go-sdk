package main

func maximalSquare1(matrix [][]byte) int {
	maxSide := 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}
	//fmt.Println(dp)
	if maxSide == 0 {
		return maxSide
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j-1]), dp[i-1][j]) + 1
				if maxSide < dp[i][j] {
					maxSide = dp[i][j]
				}
			}
		}
		//fmt.Println(dp)
	}
	return maxSide * maxSide
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
