package main

func main() {

}

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	return backtrack(nums, 0, sum/2)
}

func backtrack(nums []int, start, target int) bool {
	if target == 0 {
		return true
	}
	for i := start; i < len(nums); i++ {
		if backtrack(nums, i+1, target-nums[i]) {
			return true
		}
	}
	return false
}
