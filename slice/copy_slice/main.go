package main

import "fmt"

/*
@Time : 2020/1/10 21:41
@Author : hejun
*/

func main() {
	s := make([]int, 4, 4)
	//s = append(s, 6)
	s[0] = 6
	s1 := []int{1, 2}
	// copy方法都是从slice.data[0]位置处开始拷贝，长度取s和s1长度的最小值
	copy(s, s1)
	fmt.Println(s) //[1 2 0 0]
}
