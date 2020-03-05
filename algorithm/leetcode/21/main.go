package main

import "fmt"

/*
@Time : 2020/3/5
@Author : hejun
*/
/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

*/

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	listNode := mergeTwoLists1(l1, l2)
	for listNode != nil {
		fmt.Println(listNode.Val)
		listNode = listNode.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	curr := result
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val < l2.Val {
			curr.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
			curr = curr.Next
		} else {
			curr.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
			curr = curr.Next
		}
	}
	if l1 != nil {
		for l1 != nil {
			curr.Next = &ListNode{Val: l1.Val}
			curr = curr.Next
			l1 = l1.Next
		}
	}
	if l2 != nil {
		for l2 != nil {
			curr.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
			curr = curr.Next
		}

	}
	return result.Next
}
