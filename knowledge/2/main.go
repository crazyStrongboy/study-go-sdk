package main

import "fmt"

/*
@Time : 2020/1/3
@Author : hejun
*/

func main() {
	var pi *int = nil
	var pb *bool = nil
	var x interface{} = pi
	var y interface{} = pb
	var z interface{} = nil

	fmt.Println(x == y)   // false
	fmt.Println(x == nil) // false
	fmt.Println(y == nil) // false
	fmt.Println(x == z)   // false
	fmt.Println(y == z)   // false
}
