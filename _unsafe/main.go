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
	_    int
	Name string
}

func main() {
	a := A{Name: "hejun"}
	align := unsafe.Alignof(a)
	fmt.Printf("align: %d\n", align)
	size := unsafe.Sizeof(a)
	fmt.Printf("size: %d\n", size)
	offset := unsafe.Offsetof(a.Name)
	fmt.Printf("offset: %d,name: %s\n", offset, *(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + offset)))
}
