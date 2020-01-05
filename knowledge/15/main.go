package main

import (
	"errors"
	"fmt"
)

/*
@Time : 2020/1/5 18:30
@Author : hejun
*/
func main() {
	notfound := "not found"
	a, b := errors.New(notfound), errors.New(notfound)
	fmt.Println(a == b) // false
}
