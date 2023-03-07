package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}

func maxCoins(nums []int) int {
	t := &T{}
	t.backtrack(nums, 0)
	return t.result
}

type T struct {
	result int
}

func (t *T) backtrack(nums []int, coins int) {
	if len(nums) == 0 {
		t.result = max(t.result, coins)
		return
	}
	for i := 0; i < len(nums); i++ {
		pre := 0
		if i-1 < 0 {
			pre = 1
		} else {
			pre = nums[i-1]
		}
		next := 0
		if i+1 > len(nums)-1 {
			next = 1
		} else {
			next = nums[i+1]
		}
		dst := make([]int, 0, len(nums)-1)
		dst = append(dst, nums[:i]...)
		dst = append(dst, nums[i+1:]...)
		t.backtrack(dst, coins+pre*nums[i]*next)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
