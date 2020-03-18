package main

import "fmt"

/*
@Time : 2020/3/6
@Author : hejun
*/

/**
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。

*/

func main() {
	size := removeDuplicates2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(size)
}

func removeDuplicates(nums []int) int {
	j := 0
	for {
		if j == len(nums) {
			break
		}
		filterNum := nums[j]
		for i := j + 1; i < len(nums); i++ {
			if nums[i] == filterNum {
				nums = append(nums[0:i], nums[i+1:]...)
				i--
			} else {
				break
			}
		}
		j++
	}
	fmt.Println(nums)
	return len(nums)
}

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums)
	return i + 1
}

func removeDuplicates2(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return nums[0 : i+1]
}
