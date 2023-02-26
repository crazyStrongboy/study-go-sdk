package main

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	t := &t{}
	t.recursion(root, 1)
	return t.result
}

func (t *t) recursion(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if len(t.result) < depth {
		t.result = append(t.result, []int{})
	}
	t.result[depth-1] = append(t.result[depth-1], root.Val)
	t.recursion(root.Left, depth+1)
	t.recursion(root.Right, depth+1)
}

type t struct {
	result [][]int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
