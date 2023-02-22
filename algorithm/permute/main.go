package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var t t
	t.used = map[int]bool{}
	t.backtrack(nums, []int{})
	return t.result
}

type t struct {
	result [][]int
	used   map[int]bool
}

func (t *t) backtrack(nums, trace []int) {
	if len(trace) == len(nums) {
		r := make([]int, len(trace))
		copy(r, trace)
		t.result = append(t.result, r)
		return
	}
	for i := 0; i < len(nums); i++ {
		if t.used[nums[i]] {
			continue
		}
		trace = append(trace, nums[i])
		t.used[nums[i]] = true
		t.backtrack(nums, trace)
		trace = trace[:len(trace)-1]
		t.used[nums[i]] = false
	}
}
