package main

import (
	"fmt"
	"github.com/crazyStrongboy/study-go-sdk/algorithm/linked"
)

/*
@Time : 2020/1/17
@Author : hejun
*/

func main() {
	s := "我爱爱你"
	rs := []rune(s)
	l := linked.ConvertToList(rs)
	l.String()
	fmt.Printf("\n")
	b := palindrome(l)
	fmt.Println(b)
}

func palindrome(l *linked.ListNode) bool {
	if l == nil || l.Next == nil {
		return false
	}
	var fast *linked.ListNode = l
	var slow *linked.ListNode = l
	var prev *linked.ListNode
	for {
		fast = fast.Next.Next
		Next := slow.Next
		slow.Next = prev
		prev = slow
		slow = Next
		if fast == nil || fast.Next == nil {
			break
		}
	}
	if fast != nil {
		// 奇数位，slow节点向后移动一位进行比较
		slow = slow.Next
	}
	for {
		if prev.Value != slow.Value {
			return false
		}
		prev = prev.Next
		slow = slow.Next
		if slow == nil {
			break
		}
	}
	return true
}
