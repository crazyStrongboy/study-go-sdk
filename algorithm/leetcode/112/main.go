package main

/*
@Time : 2020/3/10
@Author : hejun
*/

/**
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil {
		root.Left.Val = root.Val + root.Left.Val
		if root.Left.Val == sum && root.Left == nil && root.Right == nil {
			return true
		}
	}
	if root.Right != nil {
		root.Right.Val = root.Val + root.Right.Val
		if root.Right.Val == sum && root.Left == nil && root.Right == nil {
			return true
		}
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
