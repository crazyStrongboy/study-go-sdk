package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{3, 4, 2, 3, 2, 2}))
}

func largestRectangleArea(heights []int) int {
	var stack []int
	stack = append(stack, 0)
	var result int
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	//fmt.Println(heights)
	for i := 1; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			result = max(result, (i-left-1)*heights[mid])
			//fmt.Println("yy: ", stack, result)
		}
		stack = append(stack, i)
		//fmt.Println("xx: ", stack, result)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
