package main

import "fmt"

func main() {
	exist := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Println(exist, "ABCCED")
}

func exist(board [][]byte, word string) bool {
	t := &t{
		used: make(map[int]int),
	}
	for i := 0; i < len(board); i++ {
		t.used[i] = -1
	}
	w := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == w[0] {
				if t.backtrack(board, word, 1, i, j) {
					return true
				}
			}
		}
	}
	return false
}

type t struct {
	used map[int]int
}

func (t *t) backtrack(board [][]byte, word string, index, x, y int) bool {
	if t.used[x] == y {
		return false
	}
	t.used[x] = y
	defer func() {
		t.used[x] = -1
	}()
	m := len(board)
	n := len(board[0])
	w := []byte(word)
	if len(w)-1 != index {
		if x-1 >= 0 && y-1 >= 0 && board[x-1][y-1] == w[index] {
			return t.backtrack(board, word, index+1, x-1, y-1)
		}
		if x-1 >= 0 && y < n && board[x-1][y] == w[index] {
			return t.backtrack(board, word, index+1, x-1, y)
		}
		if x+1 < m && y+1 < n && board[x+1][y+1] == w[index] {
			return t.backtrack(board, word, index+1, x+1, y+1)
		}
		if x < m && y+1 < n && board[x][y+1] == w[index] {
			return t.backtrack(board, word, index+1, x, y+1)
		}
	} else {
		if x-1 >= 0 && y-1 >= 0 && board[x-1][y-1] == w[index] {
			return true
		}
		if x-1 >= 0 && y < n && board[x-1][y] == w[index] {
			return true
		}
		if x+1 < m && y+1 < n && board[x+1][y+1] == w[index] {
			return true
		}
		if x < m && y+1 < n && board[x][y+1] == w[index] {
			return true
		}
	}

	return false
}
