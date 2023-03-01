package depth

import "github.com/crazyStrongboy/study-go-sdk/algorithm/tree"

func minDepth(root *tree.Node) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil && root.Left != nil {
		return 1 + minDepth(root.Left)
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
