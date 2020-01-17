package linked

import "fmt"

/*
@Time : 2020/1/17
@Author : hejun
*/

type ListNode struct {
	Value string
	Next  *ListNode
}

func (l *ListNode) String() {
	pl := l
	for {
		fmt.Printf(pl.Value)
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
