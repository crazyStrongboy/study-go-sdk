package main

import "fmt"

/*
@Time : 2020/3/27 23:41
@Author : hejun
*/

func main() {
	//bytes:= []byte{0,0,0,0,0,0,0,0}
	bytes := []byte{255, 255, 255, 255, 255, 255, 255, 255}
	aCache := uint64(0)
	aCache |= uint64(bytes[0])
	aCache |= uint64(bytes[1]) << (1 * 8)
	aCache |= uint64(bytes[2]) << (2 * 8)
	aCache |= uint64(bytes[3]) << (3 * 8)
	aCache |= uint64(bytes[4]) << (4 * 8)
	aCache |= uint64(bytes[5]) << (5 * 8)
	aCache |= uint64(bytes[6]) << (6 * 8)
	aCache |= uint64(bytes[7]) << (7 * 8)
	a := ^aCache
	fmt.Println(a)
}
