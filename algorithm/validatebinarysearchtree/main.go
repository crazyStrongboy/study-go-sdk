package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isValidBST(&TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 4},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 8},
		},
	}))
}

func isValidBST(root *TreeNode) bool {
	return isBST(root, math.MinInt64, math.MaxInt64)
	t := &t{pre: math.MinInt64}
	return t.inorder(root)
}

type t struct {
	pre int
}

func (t *t) inorder(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := t.inorder(root.Left)
	if root.Val <= t.pre {
		return false
	}
	t.pre = root.Val
	r := t.inorder(root.Right)
	return l && r
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBST(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}
	if root.Val <= left || root.Val >= right {
		return false
	}
	return isBST(root.Left, left, root.Val) && isBST(root.Right, root.Val, right)
}
