package main

import "fmt"

func longestValidParentheses(s string) int {
	stack := []int{-1}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[0 : len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				result = max(result, i-stack[len(stack)-1])
			}
		}
		fmt.Println(stack, result)
	}
	return result
}
