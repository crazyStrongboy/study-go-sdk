package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(combinationSum(nums, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	var track []int
	r := &r{}
	r.backtrack(candidates, track, target, 0)
	return r.result
}

type r struct {
	result [][]int
}

func (r *r) backtrack(nums []int, track []int, target, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		t := make([]int, len(track))
		copy(t, track)
		r.result = append(r.result, t)
		return
	}
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		target -= nums[i]
		r.backtrack(nums, track, target, i)
		target += nums[i]
		track = track[:len(track)-1]
	}
}
