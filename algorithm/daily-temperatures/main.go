package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
}

func dailyTemperatures(temperatures []int) []int {
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
