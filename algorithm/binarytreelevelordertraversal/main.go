package main

func main() {
	levelOrder2(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	})
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

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		queue  []*TreeNode
		result [][]int
	)
	queue = append(queue, root)
	for len(queue) > 0 {
		layer := make([]int, 0, len(queue))
		size := len(queue)
		for i := 0; i < size; i++ {
			temp := queue[0]
			queue = queue[1:]
			layer = append(layer, temp.Val)
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		result = append(result, layer)
	}

	return result
}

type t struct {
	result [][]int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
