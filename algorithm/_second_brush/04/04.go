package main

import "math"

func subsets(nums []int) [][]int {
	s := &subSet{used: make(map[int]bool)}
	s.backtrack(nums, 0, []int{})
	return s.result
}

type subSet struct {
	result [][]int
	used   map[int]bool
}

func (s *subSet) backtrack(nums []int, start int, track []int) {
	t := make([]int, len(track))
	copy(t, track)
	s.result = append(s.result, t)
	for i := start; i < len(nums); i++ {
		// if s.used[nums[i]]{
		//     continue
		// }
		track = append(track, nums[i])
		// s.used[nums[i]] = true
		s.backtrack(nums, i+1, track)
		track = track[:len(track)-1]
		// s.used[nums[i]] = false
	}
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, i, j, index int) bool {
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 || board[i][j] == '.' || board[i][j] != word[index] {
		return false
	}
	if index == len(word)-1 {
		return true
	} else {
		temp := board[i][j]
		board[i][j] = '.'
		if backtrack(board, word, i, j-1, index+1) || backtrack(board, word, i, j+1, index+1) ||
			backtrack(board, word, i-1, j, index+1) || backtrack(board, word, i+1, j, index+1) {
			return true
		}
		board[i][j] = temp
	}
	return false
}

func largestRectangleArea(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	var stack []int
	stack = append(stack, 0)
	result := 0
	for i := 1; i < len(heights); i++ {
		for heights[i] < heights[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			w := i - left - 1
			h := heights[mid]
			result = max(result, w*h)
		}
		stack = append(stack, i)
		//fmt.Println(stack,result)
	}
	return result
}

func maximalRectangle(matrix [][]byte) int {
	heights := make([]int, len(matrix[0]))
	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		//fmt.Println(heights)
		result = max(result, largestRectangleArea(heights))
	}
	return result
}

func inorderTraversal(root *TreeNode) []int {
	o := &orderTree{}
	o.traversal(root)
	return o.result
}

type orderTree struct {
	result []int
}

func (o *orderTree) traversal(root *TreeNode) {
	if root == nil {
		return
	}
	o.traversal(root.Left)
	o.result = append(o.result, root.Val)
	o.traversal(root.Right)
}

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if r != nil {
			if r.Right != nil {
				stack = append(stack, r.Right)
			}
			stack = append(stack, r)
			stack = append(stack, nil)
			if r.Left != nil {
				stack = append(stack, r.Left)
			}
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, temp.Val)
		}
	}
	return result
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}

func isValidBST(root *TreeNode) bool {
	left := math.MinInt64
	right := math.MaxInt64
	return isValid(root.Left, left, root.Val) && isValid(root.Right, root.Val, right)
}

func isValid(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}
	if root.Val <= left || root.Val >= right {
		return false
	}
	return isValid(root.Left, left, root.Val) && isValid(root.Right, root.Val, right)
}

func isSymmetric(root *TreeNode) bool {
	return equal(root.Left, root.Right)
}

func equal(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil && left.Val == right.Val {
		return equal(left.Left, right.Right) && equal(left.Right, right.Left)
	}
	return false
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		var temp []int
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			temp = append(temp, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		result = append(result, temp)
	}
	return result
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	v := preorder[0]
	root := &TreeNode{Val: v}
	index := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == v {
			index = i
			break
		}
	}
	leftInorder := inorder[0:index]
	root.Left = buildTree(preorder[1:1+len(leftInorder)], leftInorder)
	root.Right = buildTree(preorder[1+len(leftInorder):], inorder[index+1:])
	return root
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	dummy := &TreeNode{}
	head := dummy
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if r != nil {

			if r.Right != nil {
				stack = append(stack, r.Right)
			}
			if r.Left != nil {
				stack = append(stack, r.Left)
			}
			stack = append(stack, r)
			stack = append(stack, nil)
		} else {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//fmt.Println(temp.Val)
			head.Right = &TreeNode{Val: temp.Val}
			head = head.Right
		}
	}
	root.Right = dummy.Right.Right
	root.Left = nil
}

func flatten1(root *TreeNode) {
	f := &Flatten{}
	f.traverse(root)
}

type Flatten struct {
	pre *TreeNode
}

func (f *Flatten) traverse(root *TreeNode) {
	if root == nil {
		return
	}
	f.traverse(root.Right)
	f.traverse(root.Left)
	root.Right = f.pre
	root.Left = nil
	f.pre = root
	//fmt.Println(pre)
}
