package main

import "math"

/*
@Time : 2020/3/8
@Author : hejun
*/
/**
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 判断左右高度差即可
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lh := maxDepth(root.Left)
	rh := maxDepth(root.Right)
	return math.Abs(float64(lh-rh)) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
