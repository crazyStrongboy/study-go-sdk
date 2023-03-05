package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1
	for i := 0; i < len(nums); i++ {
		if prefix == 0 {
			result[i] = 0
			continue
		}
		j := i + 1
		product := prefix
		for j < len(nums) {
			product = product * nums[j]
			j++
		}
		prefix *= nums[i]
		result[i] = product
	}
	return result
}
