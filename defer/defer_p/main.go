package main

import "fmt"

/*
@Time : 2020/3/14
@Author : hejun
*/

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	arr := make([]int, 0)
	arr[1] = 1
}
