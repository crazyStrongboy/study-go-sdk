package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var (
		result []int
		stack  []*TreeNode
	)
	if root == nil {
		return 0
	}
	stack = append(stack, root)
	// 先进行层次遍历求和数组
	for len(stack) > 0 {
		size := len(stack)
		sum := 0
		for i := 1; i <= size; i++ {
			t := stack[0]
			sum += t.Val
			stack = stack[1:]
			if t.Left != nil {
				stack = append(stack, t.Left)
			}
			if t.Right != nil {
				stack = append(stack, t.Right)
			}
		}
		result = append(result, sum)
	}
	// 开始打劫
	if len(result) == 1 {
		return result[0]
	}
	if len(result) == 2 {
		return max(result[0], result[1])
	}
	dp := make([]int, len(result))
	dp[0] = result[0]
	dp[1] = max(result[0], result[1])
	for i := 2; i < len(result); i++ {
		dp[i] = max(dp[i-2]+result[i], dp[i-1])
	}
	return dp[len(result)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
