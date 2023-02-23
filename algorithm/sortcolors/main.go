package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if !big(nums[i], nums[j]) {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func big(i, j int) bool {
	if i < j {
		return true
	}
	return false
}
