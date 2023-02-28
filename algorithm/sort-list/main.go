package main

func main() {

}

func sortList(head *ListNode) *ListNode {
	t := head
	for t != nil {
		sortSingle(head)
		t = t.Next
	}
	return head
}

func sortSingle(head *ListNode) {
	min := head
	t := head
	for t != nil {
		t = t.Next
		if min.Val > t.Val {
			min = t
		}
	}
	head.Val, min.Val = min.Val, head.Val
}

type ListNode struct {
	Val  int
	Next *ListNode
}
