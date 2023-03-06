package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 3))
}

func findKthLargest(nums []int, k int) int {
	//sort.Ints(nums)
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, k int) int {
	index := partition(nums, left, right)
	if index == k {
		return nums[index]
	} else if index < k {
		return quickSelect(nums, index+1, right, k)
	}
	return quickSelect(nums, left, index-1, k)
}

func partition(nums []int, left, right int) int {
	index := rand.Intn(right-left+1) + left
	nums[index], nums[right] = nums[right], nums[index]
	count := left - 1
	for i := left; i < right; i++ {
		if nums[i] < nums[right] {
			count++
			nums[count], nums[i] = nums[i], nums[count]
		}
	}
	nums[count+1], nums[right] = nums[right], nums[count+1]
	return count + 1
}
