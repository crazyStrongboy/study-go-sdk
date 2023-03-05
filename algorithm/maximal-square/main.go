package main

import "fmt"

func main() {
	matrix := [][]byte{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 1},
	}
	fmt.Println(maximalSquare1(matrix))
}

func maximalSquare(matrix [][]byte) int {
	result := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		r := make([]int, 0, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			r = append(r, int(matrix[i][j]))
		}
		result[i] = r
	}
	fmt.Println(result)
	t := &T{
		flag: true,
	}
	depth := 1
	for depth <= len(result) {
		if !t.flag {
			break
		}
		t.traverse(result, depth)
		depth++
	}
	return t.max
}

type T struct {
	max  int
	flag bool
}

func (t *T) traverse(matrix [][]int, depth int) {
	t.flag = false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == depth {
				t.max = max(t.max, 1)
				if j+1 == len(matrix[0]) || i+1 == len(matrix) {
					break
				}
				if matrix[i][j+1] == depth &&
					matrix[i+1][j] == depth &&
					matrix[i+1][j+1] == depth {
					matrix[i][j] += 1
					t.flag = true
					t.max = (depth + 1) * (depth + 1)
				}
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
