package main

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
