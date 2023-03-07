package main

import "math"

func findUnsortedSubarray2(nums []int) int {
	min, max := math.MaxInt, math.MinInt
	left, right := -1, -1
	for i := 0; i < len(nums); i++ {
		if max > nums[i] {
			right = i
		} else {
			max = nums[i]
		}

		if min < nums[len(nums)-1-i] {
			left = len(nums) - 1 - i
		} else {
			min = nums[len(nums)-1-i]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}
