package main

import "fmt"

// todo 待学习 自底向上归并排序
func main() {
	result := sortList(&ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 10,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	})
	fmt.Println(result)
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow := head
	fast := head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}
func merge(head, tail *ListNode) *ListNode {
	if head == nil {
		return tail
	}
	if tail == nil {
		return head
	}
	if head.Val < tail.Val {
		head.Next = merge(head.Next, tail)
		return head
	}
	tail.Next = merge(head, tail.Next)
	return tail
}

type ListNode struct {
	Val  int
	Next *ListNode
}
