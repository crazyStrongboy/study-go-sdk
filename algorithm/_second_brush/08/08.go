package main

import "sort"

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	var result [][]int
	for k, v := range m {
		result = append(result, []int{k, v})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][1] > result[j][1]
	})

	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = result[i][0]
	}
	return ret
}
