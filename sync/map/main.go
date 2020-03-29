package main

import (
	"fmt"
	"sync"
)

/*
@Author : MarsJun
@Time : 2020/3/27 20:58
*/

func main() {
	m := &sync.Map{}
	m.Store(1, "aa")
	if v, ok := m.Load(1); ok {
		fmt.Println(v.(string))
	}
}
