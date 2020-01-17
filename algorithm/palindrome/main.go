package main

import "fmt"

/*
@Time : 2020/1/17
@Author : hejun
*/

type listnode struct {
	value string
	next  *listnode
}

func main() {
	s := "我爱爱你"
	rs := []rune(s)
	l := convertToList(rs)
	l.String()
	fmt.Printf("\n")
	b := palindrome(l)
	fmt.Println(b)
}

func palindrome(l *listnode) bool {
	if l == nil || l.next == nil {
		return false
	}
	var fast *listnode = l
	var slow *listnode = l
	var prev *listnode
	for {
		fast = fast.next.next
		next := slow.next
		slow.next = prev
		prev = slow
		slow = next
		if fast == nil || fast.next == nil {
			break
		}
	}
	if fast != nil {
		slow = slow.next
	}
	for {
		if prev.value != slow.value {
			return false
		}
		prev = prev.next
		slow = slow.next
		if slow == nil {
			break
		}
	}
	return true
}

// 字符串转成链表
func convertToList(rs []rune) *listnode {
	var head *listnode
	var tail *listnode
	for i, r := range rs {
		curr := &listnode{value: string(r)}
		if i == 0 {
			head = curr
			tail = curr
		} else {
			if tail != nil {
				tail.next = curr
				tail = curr
			}
		}
	}
	return head
}

func (l *listnode) String() {
	pl := l
	for {
		fmt.Printf(pl.value)
		pl = pl.next
		if pl == nil {
			break
		}
	}
}
