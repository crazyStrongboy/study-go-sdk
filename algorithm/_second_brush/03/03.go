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

func uniquePaths(m int, n int) int {
	var dp [][]int
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
		dp[i][0] = 1
	}
	for j := 1; j < n; j++ {
		dp[0][j] = 1
	}
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
		//fmt.Println(dp)
	}
	return dp[m-1][n-1]
}

func minPathSum(grid [][]int) int {
	var dp [][]int
	for i := 0; i < len(grid); i++ {
		dp = append(dp, make([]int, len(grid[0])))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs1(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, 2)
	dp[0] = 1
	dp[1] = 2
	for i := 3; i <= n; i++ {
		dp[0], dp[1] = dp[1], dp[1]+dp[0]
	}
	return dp[1]
}

func minDistance(word1 string, word2 string) int {
	var dp [][]int //dp[i][j]表示以i-1结尾，j-1结尾字符串编辑需要的最少步数
	m := len(word1)
	n := len(word2)
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] { // 两个字符相等，则不用换
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 替换一个字符  删除一个字符(插入和删除是一样的)
				dp[i][j] = min(dp[i-1][j-1]+1, min(dp[i][j-1]+1, dp[i-1][j]+1))
			}

		}
		//fmt.Println(dp)
	}
	return dp[m][n]
}

func sortColors(nums []int) {
	cur0 := 0
	cur1 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[cur1] = 1
			nums[cur0] = 0
			cur0++
			cur1++
		} else if nums[i] == 1 {
			nums[cur1] = 1
			cur1++
		}
	}
	for cur1 < len(nums) {
		nums[cur1] = 2
		cur1++
	}
}

func sortColors1(nums []int) {
	cur0 := 0
	cur1 := 0
	cur2 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[cur2] = 2
			nums[cur1] = 1
			nums[cur0] = 0
			cur0++
			cur1++
			cur2++
		} else if nums[i] == 1 {
			nums[cur2] = 2
			nums[cur1] = 1
			cur1++
			cur2++
		} else {
			nums[cur2] = 2
			cur2++
		}
	}
}

func sortColors2(nums []int) {
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if nums[p1] > nums[i] {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}
