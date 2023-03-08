package main

import (
	"fmt"
	"math"
)

func main() {
	node := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
	}

	fmt.Println(maxPathSum(node))
}

// 总共只有两种情况：
// root 作为根节点，取left+right+root.Val 比较最大值
// root 作为路径中的一个节点，则只能取左右路径中的最大值，返回max(left,right)+root.Val
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
