package main

import "runtime"

/*
@Time : 2020/3/7
@Author : hejun
*/

func main() {
	runtime.GOMAXPROCS(4)
}
