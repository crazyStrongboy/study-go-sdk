package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	fmt.Println(m)
	for i := 0; i < len(nums); i++ {
		fmt.Println(m[target-nums[i]])
		if v, ok := m[target-nums[i]]; ok && i != v {
			return []int{i, v}
		}
	}
	return []int{}
}
