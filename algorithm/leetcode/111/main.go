package main

import "fmt"

/*
@Time : 2020/3/9
@Author : hejun
*/

/**
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
*/

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	depth := minDepth(root)
	fmt.Println(depth)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}
	if root.Right != nil && root.Left == nil {
		return minDepth(root.Right) + 1
	}
	return Min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
