package main

import "fmt"

func main() {
	var r []int
	r = append(r, 1, 2, 3, 4)
	fmt.Println(r)
	copy(r[1:2], r[0:2])
	fmt.Println(r)
	r[0] = 2
	fmt.Println(r)
}
