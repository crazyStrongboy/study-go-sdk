package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for j := 0; j < len(nums); j++ {
		result[j] = 1
	}
	left, right := 1, 1
	size := len(nums)
	for i := 0; i < size; i++ {
		result[i] *= left
		left *= nums[i]
		result[size-i-1] *= right
		right *= nums[size-i-1]
	}
	return result
}

func productExceptSelf2(nums []int) []int {
	result := make([]int, len(nums))
	for j := 0; j < len(nums); j++ {
		result[j] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j != i {
				result[j] = result[j] * nums[i]
			}
		}
	}
	return result
}

// 超时
func productExceptSelf1(nums []int) []int {
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
