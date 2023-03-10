package main

func main() {

}

func numIslands(grid [][]byte) int {
	sum := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' {
				effect(grid, i, j)
				sum++
			}
		}
	}
	return sum
}

func effect(grid [][]byte, i int, j int) {
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	effect(grid, i-1, j)
	effect(grid, i, j+1)
	effect(grid, i, j-1)
	effect(grid, i+1, j)
}
