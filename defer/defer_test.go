package _defer

import (
	"fmt"
	"testing"
)

/*
@Time : 2020/3/2
@Author : hejun
*/
var g = 100

func Test_f1(t *testing.T) {
	a := f1()
	fmt.Printf("a: %d,g: %d\n", a, g)
}

func Test_f2(t *testing.T) {
	a := f2()
	fmt.Printf("a: %d,g: %d\n", a, g)
}

func f1() (a int) {
	defer func() {
		g = 200
	}()
	return g
}

func f2() (a int) {
	a = g
	defer func() {
		a = 200
	}()
	a = 0
	return a
}
