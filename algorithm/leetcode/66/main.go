package main

import "fmt"

/*
@Time : 2020/3/7
@Author : hejun
*/

/**
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。

*/

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
}

func plusOne(digits []int) []int {
	length := len(digits)
	i := 1
	for {
		addResult := digits[length-i] + 1
		if addResult >= 10 {
			digits[length-i] = addResult % 10
			i++
		} else {
			digits[length-i] = addResult
			break
		}
		if length-i < 0 {
			// 最高位补1
			//digits = append(digits, 0)
			//for i := length; i >= 1; i-- {
			//	digits[i] = digits[i-1]
			//}
			//digits[0] = 1
			digits = make([]int, length+1, length+1)
			digits[0] = 1
			break
		}
	}
	return digits
}
