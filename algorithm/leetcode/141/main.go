package main

import "fmt"

/*
@Time : 2020/3/21 23:02
@Author : hejun
*/

/**
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。



示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
*/
func main() {
	n3 := &ListNode{0, nil}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{3, n2}
	n4 := &ListNode{-4, n2}
	n3.Next = n4
	fmt.Println(hasCycle(n1))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 采用快慢指针的方式
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head.Next
	fast := head.Next.Next
	for slow != nil && fast != nil && fast.Next != nil {
		if slow.Val == fast.Val {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
