package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

// dp[i][j]  word1[0:i-1] word2[0:j-1] 相等需要的步数
// 如果word1[i-1]==word2[j-1] 那么  dp[i][j] = dp[i-1][j-1]
// 如果word1[i-1]!=word2[j-1],word1 新增一个字符，使得word1[i]==word2[j-1]  dp[i][j] =  dp[i][j-1]+1
// 如果word1[i-1]!=word2[j-1],word2 新增一个字符，使得word1[i-1]==word2[j]  dp[i][j] =  dp[i-1][j]+1
// 如果word1[i-1]!=word2[j-1],word1 替换一个字符，使得word1[i-1]==word2[j-1] dp[i][j] = dp[i-1][j-1]+1
func minDistance(word1 string, word2 string) int {
	var dp [][]int
	for i := 0; i <= len(word1); i++ {
		dp = append(dp, make([]int, len(word2)+1))
	}
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1]+1, min(dp[i][j-1]+1, dp[i-1][j]+1))
			}
		}
		//for _, ints := range dp {
		//	fmt.Println(ints)
		//}
		//fmt.Println()
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
