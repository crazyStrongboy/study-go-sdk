package main

import "fmt"

/*
@Time : 2020/1/12 22:12
@Author : hejun
*/

func main() {
	a := new(int) //runtime.newobject==>runtime.mallocgc
	fmt.Println(a)
}
