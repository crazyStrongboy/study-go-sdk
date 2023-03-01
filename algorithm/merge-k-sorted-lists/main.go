package main

import "fmt"

func main() {
	h := merge(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  10,
				Next: nil,
			},
		},
	}, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	})

	fmt.Println(h)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return merge(lists[0], lists[1])
	}
	return merge(mergeKLists(lists[0:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}

func merge(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	head := &ListNode{}
	tmp := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			tmp.Next = &ListNode{Val: left.Val}
			left = left.Next
		} else {
			tmp.Next = &ListNode{Val: right.Val}
			right = right.Next
		}
		tmp = tmp.Next
	}
	if left == nil {
		tmp.Next = right
	} else {
		tmp.Next = left
	}
	return head.Next
}
