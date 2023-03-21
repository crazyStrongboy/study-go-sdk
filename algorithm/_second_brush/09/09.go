package main

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
