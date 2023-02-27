package main

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 有环
			index1 := head
			index2 := fast
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index2
		}

	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
