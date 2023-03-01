package main

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == q || root == p || root == nil {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
