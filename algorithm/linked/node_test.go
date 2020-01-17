package linked

import (
	"fmt"
	"testing"
)

/*
@Time : 2020/1/17 22:10
@Author : hejun
*/
var l *LinkedList

func init() {
	n5 := &ListNode{Value: "5"}
	n4 := &ListNode{Value: "4", Next: n5}
	n3 := &ListNode{Value: "3", Next: n4}
	n2 := &ListNode{Value: "2", Next: n3}
	n1 := &ListNode{Value: "1", Next: n2}
	l = &LinkedList{Head: &ListNode{Next: n1}}
}

func TestLinkedList_Reverse(t *testing.T) {
	l.Reverse()
	l.Head.String()
}

func TestLinkedList_HasCycle(t *testing.T) {
	n1 := &ListNode{Value: "1"}
	n3 := &ListNode{Value: "3", Next: n1}
	//n2 := &ListNode{Value: "2", Next: n3}
	n1.Next = n3
	l = &LinkedList{Head: &ListNode{Next: n1}}
	isCycle := l.HasCycle()
	fmt.Println("isCycle: ", isCycle)
}
