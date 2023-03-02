package main

func main() {

}

type MinStack struct {
	head *node
}

type node struct {
	val  int
	min  int
	next *node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = &node{
			val:  val,
			min:  val,
			next: nil,
		}
	} else {
		node := &node{
			val:  val,
			min:  min(val, this.head.min),
			next: this.head,
		}
		this.head = node
	}

}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.val
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
