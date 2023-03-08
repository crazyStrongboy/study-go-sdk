package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

func longestValidParentheses(s string) int {
	result := 0
	for i := 1; i <= len(s); i++ {
		for j := 0; j <= i; j++ {
			if isValid(s[j:i]) {
				result = max(result, i-j)
			}
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, ')')
			continue
		}
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != s[i] {
			return false
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
