package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findUnsortedSubarray([]int{0}))
}

// 保证两端最大
func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var dp [][]int
	for j := 0; j < len(nums); j++ {
		dp = append(dp, make([]int, len(nums)))
	}
	var (
		max int = math.MinInt
		mi  []int
	)
	fmt.Println(nums)
	for j := 0; j <= len(nums)-1; j++ {
		if max < nums[j] {
			max = nums[j]
		}
		mi = append(mi, max)
	}

	fmt.Println("mi: ", mi)
	var (
		min int = math.MaxInt
		mx      = make([]int, len(nums))
	)
	for i := len(nums) - 1; i >= 0; i-- {
		if min > nums[i] {
			min = nums[i]
		}
		mx[i] = min
	}
	fmt.Println("max:", mx)

	dp[len(nums)-1][0] = 1
	result := 10001
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if j+1 <= i && nums[j+1] >= mi[j] && mx[j+1] >= nums[j] && dp[i][j] == 1 {
				dp[i][j+1] = dp[i][j]
			}
			if i-1 >= 0 && nums[i-1] <= mx[i] && mi[i-1] <= nums[i] && dp[i][j] == 1 {
				dp[i-1][j] = dp[i][j]
			}
		}
		for _, ints := range dp {
			fmt.Println(ints)
		}
		fmt.Println()
	}
	for i, xx := range dp {
		for j, _ := range xx {
			if dp[i][j] == 1 {
				result = minF(result, i-j+1)
			}
		}
	}
	return result
}

func minF(a, b int) int {
	if a < b {
		return a
	}
	return b
}
