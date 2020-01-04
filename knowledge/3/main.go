package main

import "fmt"

/*
@Time : 2020/1/4 19:26
@Author : hejun
*/

type T struct {
	Name string
}

func (t *T) Change() {
	t.Name = "change"
}

func main() {
	t := T{Name: "init"}
	fmt.Println("before: ", t.Name)
	change(t)
	fmt.Println("after: ", t.Name)
}

func change(t T) {
	t.Change()
}
