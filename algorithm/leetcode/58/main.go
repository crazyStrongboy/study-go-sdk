package main

import (
	"fmt"
	"strings"
)

/*
@Time : 2020/3/7
@Author : hejun
*/

/**

给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。



示例:

输入: "Hello World"
输出: 5
*/

func main() {
	fmt.Println(lengthOfLastWord1("b   a    "))
}

func lengthOfLastWord(s string) int {
	sArr := strings.Split(s, " ")
	for i := len(sArr) - 1; i >= 0; i-- {
		if sArr[i] != "" {
			return len(sArr[i])
		}
	}
	return 0
}

func lengthOfLastWord1(s string) int {
	end := len(s) - 1
	for end > 0 && s[end] == ' ' {
		end--
	}
	if end < 0 {
		return 0
	}
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	return end - start
}
