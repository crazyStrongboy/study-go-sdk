package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}

type ListNode struct {
	Val  int
	Next *ListNode
}
