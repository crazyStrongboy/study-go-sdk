package main

import "fmt"

func main() {
	r := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(r)
	rotate(r)
	fmt.Println(r)
}

func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < len(matrix) && i < l/2; i++ {
		for j, _ := range matrix[i] {
			matrix[i][j], matrix[l-i-1][j] = matrix[l-i-1][j], matrix[i][j]
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j, _ := range matrix[i] {
			if j <= i {
				continue
			}
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
