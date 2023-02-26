package main

import "fmt"

func main() {
	//fmt.Println("leetcode"[0:4])
	fmt.Println(wordBreak2("leetcode", []string{"leet", "code"}))
}

// wordBreak 动态规划
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

// wordBreak2 递归
func wordBreak2(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, w := range wordDict {
		m[w] = true
	}
	t := &T{
		m: m,
	}
	return t.backtrack(s, 0)
}

type T struct {
	m map[string]bool
}

func (t *T) backtrack(s string, start int) bool {
	if start >= len(s) {
		return true
	}
	for i := start; i <= len(s); i++ {
		if t.m[s[start:i+1]] && t.backtrack(s, i+1) {
			return true
		}
	}
	return false
}
