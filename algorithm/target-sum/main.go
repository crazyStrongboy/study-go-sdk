package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	t := &T{}
	if (sum-target)%2 != 0 {
		return 0
	}
	t.backtrack(nums, 0, (sum-target)/2)
	return t.incr
}

type T struct {
	incr int
}

func (t *T) backtrack(nums []int, start, target int) {
	if target == 0 {
		//fmt.Println(start)
		t.incr++
	}
	for i := start; i < len(nums); i++ {
		t.backtrack(nums, i+1, target-nums[i])
	}
}
