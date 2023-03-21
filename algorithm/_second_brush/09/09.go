package main

import (
	"math"
	"sort"
)

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		index := (nums[i] + n - 1) % n
		nums[index] += n
	}
	//fmt.Println(nums)
	var result []int
	for i := 0; i < len(nums); i++ {
		if nums[i] <= n {
			result = append(result, i+1)
		}
	}
	return result
}

func hammingDistance(x int, y int) int {
	x ^= y
	count := 0
	for x > 0 {
		count++
		x &= x - 1
	}
	return count
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum < int(math.Abs(float64(target))) {
		return 0
	}
	sum += target
	if sum < 0 {
		sum = -sum
	}
	if sum%2 != 0 {
		return 0
	}
	dp := make([]int, sum/2+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := sum / 2; j >= 0; j-- {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[sum/2]
}
func convertBST(root *TreeNode) *TreeNode {
	b := &BST{}
	b.traverse(root)
	return root
}

type BST struct {
	pre int
}

func (b *BST) traverse(root *TreeNode) {
	if root == nil {
		return
	}
	b.traverse(root.Right)
	b.pre += root.Val
	root.Val = b.pre
	b.traverse(root.Left)
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(max(depth(root.Left)+depth(root.Right), diameterOfBinaryTree(root.Left)),
		diameterOfBinaryTree(root.Right))
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func subarraySum2(nums []int, k int) int {
	pre := 0
	count := 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if v, ok := m[pre-k]; ok {
			count += v
		}
		m[pre] += 1
	}
	return count
}

func findUnsortedSubarray(nums []int) int {
	dst := make([]int, len(nums))
	copy(dst, nums)
	sort.Ints(dst)
	left := -1
	right := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[i] != dst[i] {
			left = i
			break
		}
	}
	if left == -1 {
		return 0
	}
	//fmt.Println(dst)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != dst[i] {
			right = i
			break
		}
	}

	//fmt.Println(left,right)
	return right - left + 1
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
	}
}
