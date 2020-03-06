package main

import "fmt"

/*
@Time : 2020/3/6
@Author : hejun
*/

/**
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
*/
func main() {
	i := strStr("a", "a")
	fmt.Println(i)
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
outer:
	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j := 0; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				continue outer
			}
		}
		return i
	}
	return -1
}
