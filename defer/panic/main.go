package main

import "fmt"

/*
@Time : 2020/3/15
@Author : hejun
*/

func main() {
	defer func() {
		func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

	}()

	//defer recover()
	//
	panic("xxxx")
	//f()
}

func f() {
	defer catch("f")

	g()
}

func catch(funcname string) {
	if r := recover(); r != nil {
		fmt.Println(funcname, "recover:", r)
	}
}

func g() {
	defer m()

	panic("g panic")
}

func m() {
	defer catch("m")

	panic("m panic")
}
