package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarraySum([]int{1, 2, 1, 2, 1}, 3))
}

func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		j := i
		for j < len(nums) {
			sum += nums[j]
			if sum == k {
				count++
				//fmt.Println(i, j)
			}
			j++
		}
	}
	return count
}

type T struct {
	sum int
}

var track []int

func (t *T) backtrack(nums []int, start int, k int) {
	if k == 0 {
		fmt.Println(track)
		t.sum++
	}
	for i := start; i < len(nums); i++ {
		track = append(track, i)
		t.backtrack(nums, start+1, k-nums[i])
		track = track[:len(track)-1]
	}
}
