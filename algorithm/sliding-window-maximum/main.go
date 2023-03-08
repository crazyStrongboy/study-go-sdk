package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	var (
		stack  []int
		result []int
	)
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if i-k+1 >= 0 {
			if stack[0] == i-k {
				stack = stack[1:]
			}
			result = append(result, nums[stack[0]])
		}
	}
	if k > len(nums) {
		result = append(result, nums[stack[0]])
	}
	return result
}
