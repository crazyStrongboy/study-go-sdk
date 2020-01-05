package main

import "fmt"

/*
@Time : 2020/1/5 11:38
@Author : hejun
*/

type A int
type B int

func (b B) M(x int) string {
	return fmt.Sprint(b, ": ", x)
}

func check(v interface{}) bool {
	_, has := v.(interface{ M(int) string })
	return has
}

func main() {
	var a A = 123
	var b B = 789
	fmt.Println(check(a)) // false
	fmt.Println(check(b)) // true
}
