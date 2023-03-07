package main

import "sort"

func findUnsortedSubarray(nums []int) int {
	dst := make([]int, len(nums))
	copy(dst, nums)
	sort.Ints(dst)
	left := 0
	right := len(nums) - 1
	for left <= right && nums[left] == dst[left] {
		left++
	}
	if left > right {
		return 0
	}
	for right >= 0 && nums[right] == dst[right] {
		right--
	}

	return right - left + 1
}
