package main

func invertTree(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	traverse(root.Left)
	traverse(root.Right)
}

func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head
	var pre *ListNode
	count := 0
	for fast != nil {

		fast = fast.Next
		cur := slow.Next

		slow.Next = pre
		pre = slow

		slow = cur
		count++
		if fast != nil {
			count++
			fast = fast.Next
		}
	}

	// fmt.Println(count)
	//return true
	if count%2 == 1 {
		pre = pre.Next
	}

	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}

	return true
}
