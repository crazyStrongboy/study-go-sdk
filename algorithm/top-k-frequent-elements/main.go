package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	//fmt.Println(m)
	y := &elems{}
	heap.Init(y)
	for t, v := range m {
		heap.Push(y, &elem{
			val:   t,
			count: v,
		})
	}
	result := make([]int, 0, k)
	size := y.Len()
	e := heap.Pop(y)
	for e != nil {
		if size <= k {
			result = append(result, e.(*elem).val)
		}
		size--
		if size == 0 {
			break
		}
		e = heap.Pop(y)
	}
	return result
}

type elem struct {
	val, count int
}

type elems []*elem

func (y *elems) Push(x any) {
	*y = append(*y, x.(*elem))
}

func (y *elems) Pop() any {
	old := *y
	e := old[y.Len()-1]
	*y = old[:y.Len()-1]
	return e
}

func (y elems) Len() int {
	return len(y)
}

func (y elems) Less(i, j int) bool {
	if y[i].count < y[j].count {
		return true
	}
	return false
}

func (y elems) Swap(i, j int) {
	y[i], y[j] = y[j], y[i]
}
