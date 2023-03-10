package main

import "fmt"

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

func trap1(height []int) int {
	sum := 0
	var (
		left  = make([]int, len(height))
		right = make([]int, len(height))
	)
	max := -1
	for i := 0; i < len(height); i++ {
		if max < height[i] {
			max = height[i]
		}
		left[i] = max
	}
	max = -1
	for i := len(height) - 1; i >= 0; i-- {
		if max < height[i] {
			max = height[i]
		}
		right[i] = max
	}
	//fmt.Println(left)
	//fmt.Println(right)
	for i := 1; i < len(height)-1; i++ {
		m := min(left[i], right[i])
		sum += m - height[i]
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
