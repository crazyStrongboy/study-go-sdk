package main

import "math"

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
