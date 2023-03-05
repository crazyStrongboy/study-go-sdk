package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarraySum([]int{1, 2, 1, 2, 1}, 3))
}

func subarraySum1(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		j := i
		for j < len(nums) {
			sum += nums[j]
			if sum == k {
				count++
				//fmt.Println(i, j)
			}
			j++
		}
	}
	return count
}

// nums[i]-nums[j] = k  (0<=j<=i)
// nums[j] = nums[i] - k
// 存储所有的前缀和在一个map里面，只要满足map[nums[i] - k]不为空，则表示有对应的index
func subarraySum(nums []int, k int) int {
	pre := 0
	count := 0
	m := make(map[int]int)
	m[0] = 1 // 初始化0为1，有可能是从index为0开始的
	for i := 0; i < len(nums); i++ {
		pre += nums[i] // 以i结尾的前缀和
		if v, ok := m[pre-k]; ok {
			count += v
		}
		m[pre] = m[pre] + 1
		//fmt.Println(m)
	}
	return count
}

type T struct {
	sum int
}

var track []int

func (t *T) backtrack(nums []int, start int, k int) {
	if k == 0 {
		fmt.Println(track)
		t.sum++
	}
	for i := start; i < len(nums); i++ {
		track = append(track, i)
		t.backtrack(nums, start+1, k-nums[i])
		track = track[:len(track)-1]
	}
}
