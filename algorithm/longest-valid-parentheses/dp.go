package main

// dp[i] 已i结尾的最大有效括号长度
// s[i] 必须为')'
// 分两种情况
// 1. s[i-1] == '('  dp[i] = dp[i-2]+2
// 2. s[i-1] == ')'  那么s[i-dp[i-1]-1] == '(',dp[i] =dp[i-1]+2+dp[i-dp[i-1]-2]
func longestValidParentheses2(s string) int {
	dp := make([]int, len(s))
	result := 0
	for i := 1; i < len(s); i++ {
		if s[i] != ')' {
			continue
		}
		if s[i-1] == '(' {
			dp[i] = 2
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			}
		} else if s[i-1] == ')' && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
		}
		result = max(result, dp[i])
	}
	return result
}
