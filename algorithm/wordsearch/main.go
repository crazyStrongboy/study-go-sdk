package main

import "fmt"

func main() {
	x := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Println(exist(x, "ABCB"))
}

func exist(board [][]byte, word string) bool {
	t := &t{
		m: len(board),
		n: len(board[0]),
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if t.backtrack(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

type t struct {
	used map[string]int
	m, n int
}

func (t *t) backtrack(board [][]byte, word string, index, x, y int) bool {
	w := []byte(word)
	if x < 0 || x > t.m-1 || y < 0 || y > t.n-1 || board[x][y] != w[index] || board[x][y] == '.' {
		return false
	}
	if len(w)-1 == index {
		return true
	} else {
		temp := board[x][y]
		board[x][y] = '.'
		b := t.backtrack(board, word, index+1, x, y-1) ||
			t.backtrack(board, word, index+1, x-1, y) ||
			t.backtrack(board, word, index+1, x+1, y) ||
			t.backtrack(board, word, index+1, x, y+1)
		board[x][y] = temp
		return b

	}
}
