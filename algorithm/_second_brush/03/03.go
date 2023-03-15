package main

import "sort"

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	index := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if index == -1 {
		return []int{-1, -1}
	}
	start := index
	//fmt.Println(index)
	for start >= 0 && nums[start] == target {
		start--
		//fmt.Println(start)
	}
	end := index
	for end <= len(nums)-1 && nums[end] == target {
		end++
	}
	return []int{start + 1, end - 1}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	s := &Sum{}
	s.backtrack(candidates, 0, target, []int{})
	return s.result
}

type Sum struct {
	result [][]int
}

func (s *Sum) backtrack(candidates []int, start, target int, track []int) {
	if target == 0 {
		t := make([]int, len(track))
		copy(t, track)
		s.result = append(s.result, t)
		return
	}
	for i := start; i < len(candidates); i++ {
		if target < candidates[i] {
			return
		}
		track = append(track, candidates[i])
		s.backtrack(candidates, i, target-candidates[i], track)
		track = track[:len(track)-1]
	}
}

func trap(height []int) int {
	var stack []int
	stack = append(stack, 0)
	result := 0
	for i := 1; i < len(height); i++ {
		for height[i] > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				w := (i - left - 1)
				h := min(height[i], height[left]) - height[mid]
				result += w * h
				//fmt.Println(result,left,mid,i)
			} else {
				break
			}
		}
		stack = append(stack, i)
		//fmt.Println(stack)
	}
	return result
}

func permute(nums []int) [][]int {
	p := &p{used: [21]int{}}
	p.backtrack(nums, []int{})
	return p.result
}

type p struct {
	result [][]int
	used   [21]int
}

func (p *p) backtrack(nums []int, trace []int) {
	if len(trace) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, trace)
		p.result = append(p.result, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if p.used[nums[i]+10] == 1 {
			continue
		}
		trace = append(trace, nums[i])
		p.used[nums[i]+10]++
		p.backtrack(nums, trace)
		p.used[nums[i]+10]--
		trace = trace[:len(trace)-1]
	}
}
