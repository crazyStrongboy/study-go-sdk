package main

import (
	"fmt"
	"strconv"
	"strings"
)

func invertTree(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	traverse(root.Left)
	traverse(root.Right)
}

func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head
	var pre *ListNode
	count := 0
	for fast != nil {

		fast = fast.Next
		cur := slow.Next

		slow.Next = pre
		pre = slow

		slow = cur
		count++
		if fast != nil {
			count++
			fast = fast.Next
		}
	}

	// fmt.Println(count)
	//return true
	if count%2 == 1 {
		pre = pre.Next
	}

	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}

	return true
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l == nil {
		return r
	}
	return l
}

func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	pre[0] = 1
	for i := 0; i < len(nums)-1; i++ {
		pre[i+1] = pre[i] * nums[i]
	}
	//fmt.Println(pre)
	suffix := make([]int, len(nums))
	suffix[len(suffix)-1] = 1
	for i := len(nums) - 1; i > 0; i-- {
		suffix[i-1] = suffix[i] * nums[i]
	}
	//fmt.Println(suffix)

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = pre[i] * suffix[i]
	}
	return result
}

func maxSlidingWindow(nums []int, k int) []int {
	var result, stack []int
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		if i >= k-1 {
			result = append(result, nums[stack[0]])
			if stack[0] == i-k+1 {
				stack = stack[1:]
			}
		}
	}
	return result
}

func searchMatrix(matrix [][]int, target int) bool {
	i := len(matrix) - 1
	j := 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			j++
		} else {
			i--
		}
	}
	return false
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i
	}
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i >= j*j {
				dp[i] = min(dp[i-j*j]+1, dp[i])
			}
		}
		//fmt.Println(dp)
	}
	return dp[n]
}

func moveZeroes(nums []int) {
	count := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i == count {
				continue
			}
			count++
			nums[count], nums[i] = nums[i], nums[count]
		}
	}
	//fmt.Println(nums,count)
	for i := count + 1; i < len(nums); i++ {
		nums[i] = 0
	}
}

func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			fast = 0
			for nums[slow] != nums[fast] {
				slow = nums[slow]
				fast = nums[fast]
			}
			return nums[slow]
		}
	}
	return -1
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var result []string
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			if top == nil {
				result = append(result, "null")
				continue
			}
			result = append(result, fmt.Sprintf("%v", top.Val))
			queue = append(queue, top.Left)
			queue = append(queue, top.Right)
		}
	}
	//fmt.Println(strings.Join(result,","))
	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, ",")
	root := &TreeNode{Val: s2I(arr[0])}
	arr = arr[1:]
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			if top != nil {
				if arr[0] == "null" {
					top.Left = nil
				} else {
					top.Left = &TreeNode{Val: s2I(arr[0])}
					queue = append(queue, top.Left)
				}
				arr = arr[1:]
				if arr[0] == "null" {
					top.Right = nil
				} else {
					top.Right = &TreeNode{Val: s2I(arr[0])}
					queue = append(queue, top.Right)
				}
				arr = arr[1:]
			}
		}

	}
	return root
}

func s2I(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	result := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
			result = max(result, dp[i])
		}
		//fmt.Println(dp)
	}
	return result
}

func removeInvalidParentheses(s string) []string {
	set := make(map[string]struct{})
	set[s] = struct{}{}
	t := &T{}
	t.recursion(set)
	return t.result
}

type T struct {
	result []string
}

func (t *T) recursion(set map[string]struct{}) map[string]struct{} {
	tmp := make(map[string]struct{})
	for k := range set {
		if isValid(k) {
			t.result = append(t.result, k)
		}
		if len(t.result) > 0 {
			continue
		}
		for i := 0; i < len(k); i++ {
			if k[i] == '(' || k[i] == ')' {
				tmp[k[:i]+k[i+1:]] = struct{}{}
			}
		}
	}
	if len(tmp) > 0 {
		return t.recursion(tmp)
	}
	return nil
}

func isValid(s string) bool {
	cnt := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
		} else if s[i] == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 3)
		dp[0][0] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1])
		dp[i][2] = dp[i-1][1]
	}
	//fmt.Println(dp)

	return max(dp[len(prices)-1][1], dp[len(prices)-1][2]) // 卖出状态和冷冻期取最大值
}

func maxCoins(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}
	for i := len(nums) - 3; i >= 0; i-- {
		for j := i + 2; j < len(nums); j++ {
			for k := i + 1; k < j; k++ {
				val := nums[i] * nums[k] * nums[j]
				dp[i][j] = max(dp[i][k]+val+dp[k][j], dp[i][j])
			}
		}
	}
	return dp[0][len(nums)-1]
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = 10001
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
		//fmt.Println(dp)
	}
	if dp[amount] == 10001 {
		return -1
	}
	return dp[amount]
}

func rob(root *TreeNode) int {
	arr := robSub(root)
	return max(arr[0], arr[1])
}

func robSub(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l := robSub(root.Left)
	r := robSub(root.Right)
	result := []int{max(l[0], l[1]) + max(r[0], r[1]), root.Val + l[0] + r[0]}
	//fmt.Println(result)
	return result
}
