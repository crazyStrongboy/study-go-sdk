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
