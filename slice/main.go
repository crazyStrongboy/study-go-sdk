package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
@Time : 2020/1/6
@Author : hejun
*/

func main() {
	s := []int{-1, 1, 2, 3, 4}
	s1 := s[1:3:3]
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	fmt.Println(*(*int)(unsafe.Pointer(sh.Data - unsafe.Sizeof(s[0]))))
}

func g(s []int) {
	s1 := append(s[:0:0], s[1:3:3]...)
	fmt.Println(s1)
}
