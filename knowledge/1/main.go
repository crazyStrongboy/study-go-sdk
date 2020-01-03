package main

import (
	"fmt"
	"unsafe"
)

/*
@Time : 2020/1/3
@Author : hejun
*/

type A struct {
	_    [0]int
	Name string
}

type B struct {
	Name string
	_    [0]int
}

func main() {
	//_ = foo.Config{[0]int{}, "name"}
	aSize := unsafe.Sizeof(A{})
	bSize := unsafe.Sizeof(B{})
	fmt.Printf("aSize: %d; bSize: %d\n", aSize, bSize)
}
