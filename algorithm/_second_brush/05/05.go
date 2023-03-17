package main

import (
	"math"
	"sort"
)

func maxProfit(prices []int) int {
	small := make([]int, len(prices))
	small[0] = prices[0]
	result := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < small[i-1] {
			small[i] = prices[i]
		} else {
			small[i] = small[i-1]
		}
	}
	//fmt.Println(small)
	big := make([]int, len(prices))
	big[len(prices)-1] = prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > big[i+1] {
			big[i] = prices[i]
		} else {
			big[i] = big[i+1]
		}
	}
	//fmt.Println(big)
	for i := 0; i < len(prices); i++ {
		result = max(result, big[i]-small[i])
	}
	return result
}

func maxProfit2(prices []int) int {
	min := math.MaxInt64
	result := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			result = max(result, prices[i]-min)
		}
	}
	return result
}

func maxPathSum(root *TreeNode) int {
	m := &maxPath{result: -1001}
	m.traverse(root)
	return m.result
}

type maxPath struct {
	result int
}

func (m *maxPath) traverse(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := m.traverse(root.Left)
	r := m.traverse(root.Right)

	if l < 0 {
		l = 0
	}
	if r < 0 {
		r = 0
	}
	m.result = max(m.result, root.Val+l+r)
	return max(l+root.Val, r+root.Val)
}

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		cur := 1

		add := nums[i] + 1
		for m[add] != 0 {
			cur += 1
			add++
		}

		sub := nums[i] - 1
		for m[sub] != 0 {
			cur += 1
			sub--
		}

		result = max(result, cur)
	}
	return result
}

func longestConsecutive1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	//fmt.Println(nums)
	dp := make([]int, len(nums))
	result := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 0 {
			dp[i] = dp[i-1]
			result = max(result, dp[i])
		}
		if nums[i]-nums[i-1] == 1 {
			dp[i] = dp[i-1] + 1
			result = max(result, dp[i])
		}

	}
	// fmt.Println(dp)
	return result + 1
}

func singleNumber(nums []int) int {
	arr := [60000]int{}
	for i := 0; i < len(nums); i++ {
		arr[nums[i]+30000]++
	}
	for i := 0; i < 60000; i++ {
		if arr[i] == 1 {
			return i - 30000
		}
	}
	return -1
}

func singleNumber1(nums []int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		ret ^= nums[i]
	}

	return ret
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	m := make(map[string]bool)
	for j := 0; j < len(wordDict); j++ {
		m[wordDict[j]] = true
	}
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if m[s[j:i]] {
				dp[i] = dp[i] || dp[j]
			}
		}
		//fmt.Println(dp)
	}
	return dp[len(s)]
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil {

		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if fast == slow {
			return true
		}

	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head.Next
	fast := head.Next.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	for head != fast {
		head = head.Next
		fast = fast.Next
	}
	return head
}

type LRUCache struct {
	m          map[int]*Node
	tail, head *Node
	capacity   int
}

type Node struct {
	Key, Val int
	Pre      *Node
	Next     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
		m:        make(map[int]*Node),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		release(v)
		insertAfterHead(this.head, v)
		return v.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		release(v)
		insertAfterHead(this.head, v)
		v.Val = value
	} else {
		node := &Node{Val: value, Key: key}
		insertAfterHead(this.head, node)
		this.m[key] = node
	}
	if this.capacity < len(this.m) {
		node := removeTail(this.tail)
		delete(this.m, node.Key)
	}
}

func release(node *Node) {
	pre := node.Pre
	next := node.Next
	pre.Next = next
	next.Pre = pre
}

func insertAfterHead(head, node *Node) {
	next := head.Next
	head.Next = node
	node.Pre = head
	node.Next = next
	next.Pre = node
}

func removeTail(tail *Node) *Node {
	node := tail.Pre
	pre := node.Pre
	pre.Next = tail
	tail.Pre = pre
	return node
}
