package main

import (
	"fmt"
	"strconv"
)

/*
@Time : 2020/3/7
@Author : hejun
*/

/**
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
*/

func main() {
	fmt.Println(addBinary("111", "110000"))
}

func addBinary(a string, b string) string {
	a_len := len(a) - 1
	b_len := len(b) - 1
	tmp := 0
	s := ""
	for a_len >= 0 && b_len >= 0 {
		at := a[a_len] - 48
		bt := b[b_len] - 48
		v := int(at) + int(bt) + tmp
		if v >= 2 {
			tmp = 1
		} else {
			tmp = 0
		}
		s = strconv.Itoa(v%2) + s
		a_len--
		b_len--
	}
	if a_len >= 0 {
		if tmp > 0 {
			for i := a_len; i >= 0; i-- {
				at := a[i] - 48
				v := int(at) + tmp
				if v >= 2 {
					tmp = 1
				} else {
					tmp = 0
				}
				s = strconv.Itoa(v%2) + s
			}
		} else {
			s = a[:a_len+1] + s
		}
	}
	if b_len >= 0 {
		if tmp > 0 {
			for i := b_len; i >= 0; i-- {
				at := b[i] - 48
				v := int(at) + tmp
				if v >= 2 {
					tmp = 1
				} else {
					tmp = 0
				}
				s = strconv.Itoa(v%2) + s
			}
		} else {
			s = b[:b_len+1] + s
		}
	}
	if tmp == 1 {
		s = strconv.Itoa(1) + s
	}
	return s
}
