package main

import "fmt"

func main() {
	fmt.Println(topKFrequent1([]int{1, 3, 3, 2, 2, 3}, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
