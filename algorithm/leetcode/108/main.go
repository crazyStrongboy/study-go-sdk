package main

import "fmt"

/*
@Time : 2020/3/8
@Author : hejun
*/

/**
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
*/

func main() {
	bst := sortedArrayToBST([]int{0, 1, 2, 3, 4, 5, 6, 7})
	fmt.Println(bst)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return arrayToBST(nums, 0, len(nums)-1)
}

func arrayToBST(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = arrayToBST(nums, left, mid-1)
	root.Right = arrayToBST(nums, mid+1, right)
	return root
}
