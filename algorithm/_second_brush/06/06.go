package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MinStack struct {
	stack []*elem
}

type elem struct {
	min, val int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, &elem{
			min: val,
			val: val,
		})
	} else {
		min := this.GetMin()
		if min > val {
			min = val
		}
		this.stack = append(this.stack, &elem{
			min: min,
			val: val,
		})
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := getListNodeLen(headA)
	b := getListNodeLen(headB)
	if a < b {
		a, b = b, a
		headA, headB = headB, headA
	}
	// fmt.Println(a,b)
	step := a - b
	for headA != nil && step > 0 {
		headA = headA.Next
		step--
	}

	for headA != nil && headB != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	if headA == headB {
		return headA
	}
	return nil
}

func getListNodeLen(head *ListNode) int {
	cnt := 0
	for head != nil {
		cnt++
		head = head.Next
	}
	return cnt
}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	dp[2] = max(nums[0], nums[1])
	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i-1], dp[i-1])
	}
	return dp[len(nums)]
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				gray(grid, i, j)
				count++
			}
		}
	}
	return count
}

func gray(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[0])-1 {
		return
	}
	if grid[i][j] == '1' {
		grid[i][j] = '2'
		gray(grid, i, j-1)
		gray(grid, i, j+1)
		gray(grid, i-1, j)
		gray(grid, i+1, j)
	}
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		cur := head.Next
		head.Next = pre
		pre = head
		head = cur
	}
	return pre
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make(map[int]map[int]bool) // 记录学了k课程可以解锁哪些课程
	depth := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		depth[i] = 0
	}
	for i := 0; i < len(prerequisites); i++ {
		require := prerequisites[i]
		if m[require[1]] == nil {
			m[require[1]] = make(map[int]bool)
		}
		m[require[1]][require[0]] = true
		depth[require[0]]++
	}

	//fmt.Println(m)
	//fmt.Println(depth)
	var queue []int
	for k, v := range depth {
		if v == 0 {
			queue = append(queue, k) // 深度为1的课程可以先学习
			//numCourses --
		}
	}

	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		for k, _ := range m[top] {
			depth[k]--
			if depth[k] == 0 { // 前置课程都学习完了，解锁当前课程
				queue = append(queue, k)
				// numCourses --
				// if 0 == numCourses{
				//     return true
				// }
			}
		}
	}

	for _, v := range depth {
		if v != 0 {
			return false
		}
	}

	return true
}

type Trie struct {
	elem  [26]*Trie
	isEnd bool
}

func Constructor1() Trie {
	return Trie{
		elem: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if cur.elem[index] == nil {
			cur.elem[index] = &Trie{
				elem: [26]*Trie{},
			}
		}
		cur = cur.elem[index]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if cur == nil || cur.elem[index] == nil {
			return false
		}
		cur = cur.elem[index]
	}
	return cur.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if cur == nil || cur.elem[index] == nil {
			return false
		}
		cur = cur.elem[index]
	}
	return true
}

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	left := 0
	right := len(nums) - 1
	mid := rand.Intn(right-left+1) + left
	fmt.Println(nums)
	nums[right], nums[mid] = nums[mid], nums[right]
	fmt.Println(nums)
	fmt.Println(nums[right], nums[mid], k)
	i := 0
	count := i - 1
	for i < right {
		if nums[i] < nums[right] {
			count++
			nums[i], nums[count] = nums[count], nums[i]
		}
		i++
	}
	index := count + 1
	nums[right], nums[index] = nums[index], nums[right]
	fmt.Println(i)
	fmt.Println(nums)
	//return 0
	if right-index == k-1 {
		return nums[index]
	} else if right-index < k-1 {
		return findKthLargest(nums[left:index+1], k-(right-index))
	}
	return findKthLargest(nums[index+1:right+1], k)
}

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	slideSize := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				slideSize = 1
			}
			dp[i][j] = int(matrix[i][j] - '0')
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if dp[i][j] != 0 {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
				slideSize = max(slideSize, dp[i][j])
			}
		}
	}

	return slideSize * slideSize
}
