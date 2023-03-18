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
