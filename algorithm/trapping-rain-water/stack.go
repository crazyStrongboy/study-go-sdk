package main

import "fmt"

func trap(height []int) int {
	var stack []int
	stack = append(stack, 0)
	var sum int
	for i := 1; i < len(height); i++ {
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				h := min(height[left], height[i]) - height[mid]
				w := i - left - 1
				sum += h * w
			}
		}
		stack = append(stack, i)
		fmt.Println(stack, sum)
	}
	return sum
}
