package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{75, 75, 75, 75, 75, 91, 75, 75, 75, 75, 75, 75, 30, 75, 75, 78, 75, 75, 75, 75, 75, 7, 79, 93, 75, 75, 15,
		75, 75, 75, 75, 75, 75, 4, 75, 75, 21, 75, 75, 19, 75, 59, 75, 76, 75, 75, 75, 75, 75, 75, 75, 33, 75, 75, 75, 58, 75, 75, 5, 75, 97, 81, 48, 42, 75, 87, 75, 75, 25, 27, 94, 20, 75, 75, 75, 29, 75, 75, 86, 67, 75, 75, 75, 65, 75, 75, 75, 75, 75, 39, 75, 56, 75, 75, 75, 75, 3, 75, 75, 75}))
}

func findDuplicate1(nums []int) int {
	n := len(nums)
	x := 0
	for i := 0; i < len(nums); i++ {
		t := (nums[i] + n) % n
		//fmt.Println(1 << t)
		if x&(1<<t) > 0 {
			return nums[i]
		}
		x = x ^ (1 << t)
		fmt.Println(x)
	}
	return 0
}

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}
