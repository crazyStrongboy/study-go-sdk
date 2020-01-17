package main

import "github.com/crazyStrongboy/study-go-sdk/algorithm/linked"

/*
@Time : 2020/1/17
@Author : hejun
*/

func main() {
	s := "abcdefg"
	rs := []rune(s)
	l := linked.ConvertToList(rs)

	var prev *linked.ListNode
	// 反转链表
	for {
		if l == nil {
			break
		}
		next := l.Next
		l.Next = prev
		prev = l
		l = next
	}

	prev.String()
}
