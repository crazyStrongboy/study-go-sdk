package main

import (
	"fmt"
)

func main() {
	r := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 1},
			},
			Right: &TreeNode{Val: 4},
		},
	}
	fmt.Println(pathSum(r, 22))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	t := &T{}
	t.traverse(root, targetSum)
	return t.r
}

type T struct {
	r int
}

func (t *T) traverse(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	t.cal(root, targetSum, root.Val)
	// fmt.Print(root.Val," ")
	t.traverse(root.Left, targetSum)
	t.traverse(root.Right, targetSum)
}

func (t *T) cal(root *TreeNode, targetSum int, result int) {
	if root == nil {
		return
	}
	if result == targetSum {
		t.r++
		// fmt.Println("root: ",root.Val)
	}
	if root.Left != nil {
		temp := result + root.Left.Val
		// if temp == targetSum {
		// 	fmt.Println("left: ",root.Left.Val)
		// 	t.r++
		// }
		t.cal(root.Left, targetSum, temp)
	}
	if root.Right != nil {
		temp := result + root.Right.Val
		// if temp == targetSum {
		// 	fmt.Println("right: ",root.Right.Val)
		// 	t.r++
		// }
		t.cal(root.Right, targetSum, temp)
	}
}
