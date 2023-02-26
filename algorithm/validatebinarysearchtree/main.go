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
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 8},
		},
	}))
}

func isValidBST(root *TreeNode) bool {
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
