package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	t := &t{}
	t.backtrack(nums, []int{}, 0)
	return t.result
}

type t struct {
	result [][]int
}

func (t *t) backtrack(nums, track []int, start int) {
	p := make([]int, len(track))
	copy(p, track)
	t.result = append(t.result, p)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		t.backtrack(nums, track, i+1)
		track = track[0 : len(track)-1]
	}
}
