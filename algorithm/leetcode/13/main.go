package main

import "fmt"

/*
@Time : 2020/3/5
@Author : hejun
*/
/**
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

func main() {
	fmt.Println(romanToInt2("IV"))
}

var mapping = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	ret := 0
	l := len(s)
	for i := 0; i < l; i++ {
		v := mapping[s[i]]
		if s[i] == 'I' && (i+1) < l && (s[i+1] == 'V' || s[i+1] == 'X') {
			v = -v
		} else if s[i] == 'X' && (i+1) < l && (s[i+1] == 'L' || s[i+1] == 'C') {
			v = -v
		} else if s[i] == 'C' && (i+1) < l && (s[i+1] == 'D' || s[i+1] == 'M') {
			v = -v
		}
		ret += v
	}
	return ret
}

func romanToInt1(s string) int {
	ret := 0
	l := len(s)
	for i := 0; i < l; i++ {
		v := mapping[s[i]]
		if i+1 < l && v < mapping[s[i+1]] {
			v = -v
		}
		ret += v
	}
	return ret
}

func romanToInt2(s string) int {
	ret := 0
	l := len(s)
	pre := mapping[s[0]]
	for i := 1; i < l; i++ {
		v := mapping[s[i]]
		if pre < v {
			ret -= pre
		} else {
			ret += pre
		}
		pre = v
	}
	ret += pre
	return ret
}
