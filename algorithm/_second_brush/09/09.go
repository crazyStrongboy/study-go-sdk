package main

import "math"

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		index := (nums[i] + n - 1) % n
		nums[index] += n
	}
	//fmt.Println(nums)
	var result []int
	for i := 0; i < len(nums); i++ {
		if nums[i] <= n {
			result = append(result, i+1)
		}
	}
	return result
}

func hammingDistance(x int, y int) int {
	x ^= y
	count := 0
	for x > 0 {
		count++
		x &= x - 1
	}
	return count
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum < int(math.Abs(float64(target))) {
		return 0
	}
	sum += target
	if sum < 0 {
		sum = -sum
	}
	if sum%2 != 0 {
		return 0
	}
	dp := make([]int, sum/2+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := sum / 2; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[sum/2]
}
