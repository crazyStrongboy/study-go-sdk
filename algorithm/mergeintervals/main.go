package main

import (
	"fmt"
	"sort"
)

func main() {

}

func merge(intervals [][]int) [][]int {

	var result [][]int

	if len(intervals) == 0 {
		return result
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)
	i := 1
	result = append(result, intervals[0])
	for i < len(intervals) {

		r := result[len(result)-1]

		temp := intervals[i]
		if r[1] >= temp[0] && r[1] <= temp[1] {
			r[1] = temp[1]
		} else if temp[1] <= r[1] {

		} else {
			result = append(result, temp)
		}
		i++
	}
	return result
}
