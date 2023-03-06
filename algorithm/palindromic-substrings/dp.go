package main

// dp[i][j]  表示[j:i] 是回文串
// 1. 如果 s[i] == s[j] ，那么dp[i][j] = dp[i-1][j+1]
// 2. 如果i==j,那么dp[i][j] = true
func countSubstrings(s string) int {
	var dp [][]bool
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			if i == j || (s[i] == s[j] && (dp[i-1][j+1] || i-j == 1)) {
				dp[i][j] = true
				sum++
			}
			//for _, bools := range dp {
			//	fmt.Println(bools)
			//}
			//
			//fmt.Println()
		}
	}
	return sum
}
