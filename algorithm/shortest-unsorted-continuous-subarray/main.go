package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findUnsortedSubarray([]int{1}))
}

// 保证两端最大
func findUnsortedSubarray1(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var (
		max int = math.MinInt
		mi  []int
	)
	//fmt.Println(nums)
	for j := 0; j <= len(nums)-1; j++ {
		if max < nums[j] {
			max = nums[j]
		}
		mi = append(mi, max)
	}

	//fmt.Println("mi: ", mi)
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
	//fmt.Println("max:", mx)

	result := len(nums)
	i := len(nums) - 1
	j := 0
	for {
		if i-1 >= 0 && nums[i-1] <= nums[i] && mi[i-1] <= nums[i] {
			result--
			i--
			continue
		} else if i == 0 {
			if nums[0] <= mx[1] {
				result--
				i--
				continue
			}
		}
		break
	}
	if result == 0 {
		return 0
	}
	for {
		if j+1 <= len(nums)-1 && nums[j+1] >= nums[j] && mx[j+1] >= nums[j] {
			result--
			j++
			continue
		} else if j == len(nums)-1 {
			if nums[len(nums)-1] >= mi[len(nums)-2] {
				result--
				i--
				continue
			}
		}
		break

	}
	return result
}
