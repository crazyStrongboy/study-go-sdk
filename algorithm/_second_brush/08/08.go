package main

import (
	"container/heap"
	"sort"
)

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	var result [][]int
	for k, v := range m {
		result = append(result, []int{k, v})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][1] > result[j][1]
	})

	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = result[i][0]
	}
	return ret
}

func topKFrequent1(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	e := elems{}
	heap.Init(&e)
	for key, val := range m {
		heap.Push(&e, &elem{val, key})
		if e.Len() > k {
			heap.Pop(&e)
		}
	}

	ret := make([]int, k)
	for i := 0; i < k; i++ {
		tmp := heap.Pop(&e)
		ret[i] = tmp.(*elem)[1]
	}

	return ret
}

type elem [2]int

type elems []*elem

func (e elems) Len() int {
	return len(e)
}

func (e elems) Less(i, j int) bool {
	return e[i][0] < e[j][0]
}

func (e elems) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *elems) Push(x any) {
	*e = append(*e, x.(*elem))
}

func (e *elems) Pop() any {
	ret := (*e)[e.Len()-1]
	*e = (*e)[:e.Len()-1]
	return ret
}

func decodeString(s string) string {
	var stack []string
	var nums []int
	var tmp []byte
	var n int
	// var result string
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stack = append(stack, string(tmp))
			nums = append(nums, n)
			n = 0
			tmp = make([]byte, 0)
		} else if s[i] == ']' {
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			size := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			for size > 0 {
				pre += string(tmp)
				size--
			}
			//result=pre
			tmp = []byte(pre)
			//fmt.Println(pre)
		} else if s[i] >= 'a' && s[i] <= 'z' {
			tmp = append(tmp, s[i])
		} else {
			n = n*10 + int(s[i]-'0')
			// nums = append(nums,int(s[i]-'0'))
		}
	}
	//result = result+string(tmp)
	return string(tmp)
}
