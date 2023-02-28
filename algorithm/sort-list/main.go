package main

func main() {

}

func sortList(head *ListNode) *ListNode {
	t := head
	for t != nil {
		sortSingle(t)
		t = t.Next
	}
	return head
}

func sortSingle(head *ListNode) {
	min := head
	t := head
	for t != nil {
		if min.Val > t.Val {
			min = t
		}
		t = t.Next
	}
	head.Val, min.Val = min.Val, head.Val
}

type ListNode struct {
	Val  int
	Next *ListNode
}
