package main

import "fmt"

func main() {

}

// 转换成求每一层柱子的最大面积
func maximalRectangle(matrix [][]byte) int {
	result := 0
	height := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		result = max(result, calRectangle(height))
		fmt.Println(result)
	}
	return result
}

func calRectangle(height []int) int {
	var stack []int
	result := 0
	height = append([]int{0}, height...)
	height = append(height, 0)
	stack = append(stack, height[0])
	for i := 1; i < len(height); i++ {
		for height[i] < height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			w := i - left - 1
			h := height[mid]
			result = max(result, w*h)
		}
		stack = append(stack, i)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
