package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct1([]int{2, 3, -2, 4}))
}

func maxProduct(nums []int) int {
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		product := nums[i]
		if max < product {
			max = product
		}
		for j := i + 1; j < len(nums); j++ {
			product = product * nums[j]
			if max < product {
				max = product
			}
		}
	}
	return max
}

func maxProduct1(nums []int) int {
	imax := math.MinInt64
	if len(nums) == 0 {
		return 0
	}

	x, y := 1, 1 // x 最大值  y 最小值
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			x, y = y, x
		}
		x = max(x*nums[i], nums[i])
		y = min(y*nums[i], nums[i])
		imax = max(imax, x)
	}
	return imax
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
