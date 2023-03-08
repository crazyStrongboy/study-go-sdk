package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures1(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		j := i + 1
		for j < len(temperatures) {
			if temperatures[j] > temperatures[i] {
				result[i] = j - i
				break
			}
			j++
		}
	}
	return result
}

// 倒着遍历，如果temperatures[i] < temperatures[j],那么result[i] = j-i
// 如果temperatures[i] > temperatures[j],并且result[j]==0,那么result[i] =0
func dailyTemperatures2(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i := len(temperatures) - 2; i >= 0; i-- {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[i] < temperatures[j] {
				result[i] = j - i
				break
			} else if result[j] == 0 {
				//result[i] = 0 // 这里可以注释掉不写
				break
			}
		}
	}
	return result
}

// 单调栈
func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var stack []int
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		fmt.Println(stack)
	}
	return result
}
