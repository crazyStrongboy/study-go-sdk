package main

import "fmt"

func main() {
	fmt.Println(wordBreak("aaaaaaa", []string{"aaa", "aaaa"}))
}

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, w := range wordDict {
		m[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if m[s[j:i]] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	fmt.Println(dp)
	return dp[len(s)]
}
