package linked

import "fmt"

/*
@Time : 2020/1/17
@Author : hejun
*/

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
}

// 链表反转
func (l *LinkedList) Reverse() {
	if l.Head == nil || l.Head.Next == nil || l.Head.Next.Next == nil {
		return
	}
	curr := l.Head.Next
	var prev *ListNode = nil
	for curr != nil {
		// 指针后移一位
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head.Next = prev
}

// 判断是否是环
func (l *LinkedList) HasCycle() bool {
	if l.Head == nil || l.Head.Next == nil || l.Head.Next.Next == nil {
		return false
	}
	faster := l.Head.Next
	slower := l.Head.Next
	for faster != nil && faster.Next != nil {
		// 快指针一次走慢指针的两倍
		faster = faster.Next.Next
		slower = slower.Next
		if faster == slower {
			return true
		}
	}
	return false
}

func (l *ListNode) String() {
	pl := l
	for {
		fmt.Printf("%+v", pl.Value)
		pl = pl.Next
		if pl == nil {
			break
		}
	}
}

// 字符串转成链表
func ConvertToList(rs []rune) *ListNode {
	var head *ListNode
	var tail *ListNode
	for i, r := range rs {
		curr := &ListNode{Value: string(r)}
		if i == 0 {
			head = curr
			tail = curr
		} else {
			if tail != nil {
				tail.Next = curr
				tail = curr
			}
		}
	}
	return head
}
