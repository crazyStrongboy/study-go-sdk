package main

func main() {

}

func searchMatrix(matrix [][]int, target int) bool {
	i := len(matrix) - 1
	j := 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] > target {
			i--
			continue
		}
		if matrix[i][j] < target {
			j++
			continue
		}
		if matrix[i][j] == target {
			return true
		}
	}
	return false
}
