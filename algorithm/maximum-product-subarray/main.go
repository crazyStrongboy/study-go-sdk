package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
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
