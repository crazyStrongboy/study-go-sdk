package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	fmt.Println(m)
	for i := 0; i < len(nums); i++ {
		fmt.Println(m[target-nums[i]])
		if v, ok := m[target-nums[i]]; ok && i != v {
			return []int{i, v}
		}
	}
	return []int{}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	pre := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + pre
		v := sum % 10
		pre = sum / 10
		head.Next = &ListNode{Val: v}
		head = head.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + pre
		v := sum % 10
		pre = sum / 10
		head.Next = &ListNode{Val: v}
		head = head.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + pre
		v := sum % 10
		pre = sum / 10
		head.Next = &ListNode{Val: v}
		head = head.Next
		l2 = l2.Next
	}

	if pre != 0 {
		head.Next = &ListNode{Val: pre}
	}

	return dummy.Next
}
