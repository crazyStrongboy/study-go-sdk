package main

import "fmt"

/*
@Time : 2020/3/5
@Author : hejun
*/

/**
排排坐，分糖果。

我们买了一些糖果 candies，打算把它们分给排好队的 n = num_people 个小朋友。

给第一个小朋友 1 颗糖果，第二个小朋友 2 颗，依此类推，直到给最后一个小朋友 n 颗糖果。

然后，我们再回到队伍的起点，给第一个小朋友 n + 1 颗糖果，第二个小朋友 n + 2 颗，依此类推，直到给最后一个小朋友 2 * n 颗糖果。

重复上述过程（每次都比上一次多给出一颗糖果，当到达队伍终点后再次从队伍起点开始），直到我们分完所有的糖果。注意，就算我们手中的剩下糖果数不够（不比前一次发出的糖果多），这些糖果也会全部发给当前的小朋友。

返回一个长度为 num_people、元素之和为 candies 的数组，以表示糖果的最终分发情况（即 ans[i] 表示第 i 个小朋友分到的糖果数）。

*/

func main() {
	candies := 10
	num_people := 3
	ret := distributeCandies1(candies, num_people)
	//ret := distributeCandies(candies, num_people)
	fmt.Println(ret)
}

// 暴力法
func distributeCandies(candies int, num_people int) []int {
	ret := make([]int, num_people, num_people)
	num := 1
	for candies > 0 {
		for i := 0; i < num_people && candies > 0; i++ {
			if candies > num {
				ret[i] += num
			} else {
				ret[i] += candies
			}
			candies -= num
			num++
		}
	}
	return ret
}

// 进行取模计算，每次分的（糖果数量-1）%总人数  =  分配的当前人的下标
func distributeCandies1(candies int, num_people int) []int {
	ret := make([]int, num_people, num_people)
	curr_give := 0
	for candies > 0 {
		index := curr_give % num_people
		curr_give++
		ret[index] += min(curr_give, candies)
		candies -= curr_give
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
