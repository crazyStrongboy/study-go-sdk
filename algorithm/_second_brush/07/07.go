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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l == nil {
		return r
	}
	return l
}

func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	pre[0] = 1
	for i := 0; i < len(nums)-1; i++ {
		pre[i+1] = pre[i] * nums[i]
	}
	//fmt.Println(pre)
	suffix := make([]int, len(nums))
	suffix[len(suffix)-1] = 1
	for i := len(nums) - 1; i > 0; i-- {
		suffix[i-1] = suffix[i] * nums[i]
	}
	//fmt.Println(suffix)

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = pre[i] * suffix[i]
	}
	return result
}
