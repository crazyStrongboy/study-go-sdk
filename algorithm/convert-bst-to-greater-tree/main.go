package main

func main() {
	node := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8}},
		},
	}

	convertBST(node)
}

// 右-》中-》左
func convertBST(root *TreeNode) *TreeNode {
	t := &T{}
	t.traverse(root)
	return root
}

type T struct {
	pre int
}

func (t *T) traverse(root *TreeNode) {
	if root == nil {
		return
	}
	t.traverse(root.Right)
	root.Val = t.pre + root.Val
	t.pre = root.Val
	//fmt.Println(root.Val)
	t.traverse(root.Left)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
