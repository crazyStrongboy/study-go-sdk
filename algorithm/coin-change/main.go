package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{2}, 3))
}

func coinChange1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	t := &T{
		amount: amount,
		min:    math.MaxInt,
	}
	t.backtrack(coins, 0, 0)
	if t.min == math.MaxInt {
		return -1
	}
	return t.min
}

type T struct {
	amount int
	track  []int
	min    int
}

func (t *T) backtrack(coins []int, start, sum int) {
	if sum == t.amount {
		t.min = min(t.min, len(t.track))
		//fmt.Println(t.track)
		return
	}
	if len(t.track) == t.amount {
		return
	}
	for i := start; i < len(coins); i++ {
		sum += coins[i]
		t.track = append(t.track, coins[i])
		t.backtrack(coins, start, sum)
		t.track = t.track[0 : len(t.track)-1]
		sum -= coins[i]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
