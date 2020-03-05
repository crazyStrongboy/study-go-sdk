package main

import "fmt"

/*
@Time : 2020/3/5
@Author : hejun
*/
/**

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
*/

func main() {
	fmt.Println(longestCommonPrefix([]string{"aa", "a"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	base := strs[0]
	j := 0
outer:
	for j < len(base) {
		for i := 1; i < len(strs); i++ {
			curr := strs[i]
			if j >= len(curr) {
				break outer
			}
			if base[j] != curr[j] {
				break outer
			}
		}
		j++
	}
	return base[:j]
}
