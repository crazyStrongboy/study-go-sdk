package main

import (
	"container/list"
	"fmt"
)

/*
@Time : 2020/3/8
@Author : hejun
*/

/**
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/

func main() {
	root := &TreeNode{3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		}}

	ints := levelOrderBottom(root)
	fmt.Println(ints)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	ret := list.New()
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		level := make([]int, 0, size)
		queue1 := make([]*TreeNode, 0)
		for i := 0; i < size; i++ {
			element := queue[i]
			level = append(level, element.Val)
			if element.Left != nil {
				queue1 = append(queue1, element.Left)
			}
			if element.Right != nil {
				queue1 = append(queue1, element.Right)
			}
		}
		queue = queue1
		ret.PushFront(level)
	}
	curr := ret.Front()
	for curr != nil {
		result = append(result, curr.Value.([]int))
		curr = curr.Next()
	}
	return result
}
