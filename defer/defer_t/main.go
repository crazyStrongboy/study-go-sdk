package main

/*
@Time : 2020/3/6
@Author : hejun
*/

func main() {
	f()
}

func f() {
	defer sum(1, 2)
}

func sum(a, b int) int {
	return a + b
}
