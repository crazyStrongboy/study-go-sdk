package _map

import (
	"fmt"
	"testing"
)

/*
@Time : 2020/3/2
@Author : hejun
*/

func Test_f(t *testing.T) {
	//m := map[int]string{} // runtime.makemap
	m := make(map[int]string, 16)
	m[1] = "test"
	fmt.Println(m)
}
