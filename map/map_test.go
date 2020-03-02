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
	m := map[int]string{} // runtime.makemap
	m[1] = "test"
	fmt.Println(m)
}
