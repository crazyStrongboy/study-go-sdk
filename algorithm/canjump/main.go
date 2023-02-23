package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 0, 0, 0}))
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	m := make(map[int]bool)
	for i := range nums {
		j := i
		if nums[j] == 0 && j != len(nums)-1 {
			m[i] = false
			j = i - 1
			for j >= 0 {
				if nums[j] > i-j {
					m[i] = true
				}
				j--
			}
		}
	}
	for _, v := range m {
		if !v {
			return false
		}
	}
	return true
}
