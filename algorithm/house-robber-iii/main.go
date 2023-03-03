package main

import "fmt"

func main() {
	t := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	fmt.Println(rob(t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	result := traverse(root)
	return max(result[0], result[1])
}

// 数组 [0]=表示打劫自己能获取的最大值  [1]=表示打劫子节点能获取的最大值
func traverse(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l := traverse(root.Left)
	r := traverse(root.Right)

	ret := []int{root.Val + l[1] + r[1], max(l[0], l[1]) + max(r[0], r[1])}
	fmt.Println(ret)
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
