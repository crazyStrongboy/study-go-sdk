package main

import (
	"fmt"
	"math"
)

func main() {
	node := &TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val:   -2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   -3,
			Left:  &TreeNode{Val: -1},
			Right: &TreeNode{Val: -2},
		},
	}

	fmt.Println(maxPathSum(node))
}

func maxPathSum(root *TreeNode) int {
	t := &T{m: math.MinInt}
	t.traverse(root)
	return t.m
}

type T struct {
	m int
}

func (t *T) traverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := max(t.traverse(root.Left), 0)
	r := max(t.traverse(root.Right), 0)
	t.m = max(l+r+root.Val, t.m)
	return max(l, r) + root.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
