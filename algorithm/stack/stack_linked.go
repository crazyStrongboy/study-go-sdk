package stack

import "fmt"

/*
@Time : 2020/1/18 22:33
@Author : hejun
*/

type ListNode struct {
	Prev  *ListNode
	Value interface{}
	Next  *ListNode
}

type LinkedListStack struct {
	Head *ListNode
	Tail *ListNode
}

func NewLinkedListStack() *LinkedListStack {
	head, tail := &ListNode{}, &ListNode{}
	head.Next = tail
	tail.Prev = head
	return &LinkedListStack{
		Head: head,
		Tail: tail,
	}
}

func (l *LinkedListStack) Push(value interface{}) {
	node := &ListNode{Value: value}
	node.Prev = l.Tail.Prev
	l.Tail.Prev.Next = node
	node.Next = l.Tail
	l.Tail.Prev = node
}

func (l *LinkedListStack) Pop() (value interface{}) {
	if l.Head.Next == l.Tail {
		return
	}
	node := l.Tail.Prev
	node.Prev.Next = l.Tail
	l.Tail.Prev = node.Prev
	value = node
	return
}

func (l *LinkedListStack) String() string {
	s := ""
	curr := l.Head.Next
	for curr != l.Tail {
		s += fmt.Sprintf("--> %+v", curr.Value)
		curr = curr.Next
	}
	return s
}
