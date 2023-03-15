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

func rotate(matrix [][]int) {
	top := 0
	bottom := len(matrix) - 1
	for top < bottom {
		matrix[top], matrix[bottom] = matrix[bottom], matrix[top]
		top++
		bottom--
	}
	//fmt.Println(matrix)
	for i := len(matrix) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	//fmt.Println(matrix)
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	//fmt.Println(m)
	for i := 0; i < len(strs); i++ {
		k := [26]int{}
		for _, v := range strs[i] {
			k[v-'a']++
		}
		if m[k] == nil {
			m[k] = []string{}
		}
		m[k] = append(m[k], strs[i])
	}
	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func maxSubArray(nums []int) int {
	pre := 0
	result := -10001
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if pre < nums[i] {
			pre = nums[i]
		}
		//fmt.Println(pre)
		result = max(result, pre)
	}
	return result
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	m := nums[0]
	i := 1
	for i <= m {
		m = max(i+nums[i], m)
		if m >= len(nums)-1 {
			return true
		}
		i++
	}
	return false
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//fmt.Println(intervals)
	var result [][]int
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if cur[0] > pre[1] {
			result = append(result, pre)
			pre = cur
		} else {
			pre[1] = max(pre[1], cur[1])
		}
	}
	result = append(result, pre)
	return result
}
