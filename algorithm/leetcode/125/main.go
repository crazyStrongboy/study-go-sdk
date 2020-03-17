package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
@Time : 2020/3/17 23:19
@Author : hejun
*/
/**
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
*/

func main() {
	fmt.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	i := 0
	j := len(s) - 1
	for i <= j {
		p := s[i]
		q := s[j]
		if !unicode.IsLetter(rune(p)) && !unicode.IsNumber(rune(p)) {
			i++
			continue
		}
		if !unicode.IsLetter(rune(q)) && !unicode.IsNumber(rune(q)) {
			j--
			continue
		}

		if strings.ToUpper(string(p)) != strings.ToUpper(string(q)) {
			return false
		}
		i++
		j--
	}
	return true
}
