package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   5,
			Right: &TreeNode{Val: 6},
		},
	}
	flatten(root)
	fmt.Println(root)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	t := &T{}
	t.recursion(root)
	root.Right = t.Head.Right
	root.Val = t.Head.Val
	root.Left = nil
}

func (t *T) recursion(root *TreeNode) {

	if root == nil {
		return
	}
	if root.Val == 5 {
		fmt.Println("111")
	}
	if t.Head == nil {
		t.Head = &TreeNode{Val: root.Val}
		t.Next = t.Head
	} else {
		t.Next.Right = &TreeNode{Val: root.Val}
		t.Next = t.Next.Right
	}
	t.recursion(root.Left)
	t.recursion(root.Right)
}

type T struct {
	Head *TreeNode
	Next *TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t TreeNode) String() string {
	return fmt.Sprintf("%v,%v,%v", t.Val, t.Left, t.Right)
}
