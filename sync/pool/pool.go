package main

import (
	"fmt"
	"sync"
)

/*
@Time : 2020/3/29
@Author : hejun
*/
func main() {
	p := &sync.Pool{}
	p.Put(1)
	fmt.Println(p.Get())
}
