package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reconstructQueue([][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}))
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	//fmt.Println(people)
	for i := 0; i < len(people); i++ {
		p := people[i]
		index := p[1]
		if index < i {
			copy(people[index+1:], append(people[index:i], people[i+1:]...))
			people[index] = p
		}
	}
	return people
}
