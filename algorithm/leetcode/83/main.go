package main

import "fmt"

/*
@Time : 2020/3/7
@Author : hejun
*/

/**

 */

func main() {
	head := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	listNode := deleteDuplicates(head)
	curr := listNode
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表是排序的
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := head
	next := head.Next
	for next != nil {
		if pre.Val == next.Val {
			pre.Next = next.Next
			next = next.Next
		} else {
			pre = next
			next = next.Next
		}
	}
	return head
}
