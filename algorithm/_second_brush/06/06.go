package main

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
