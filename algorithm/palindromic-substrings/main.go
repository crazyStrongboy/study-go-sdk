package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("aaa"))
}

func countSubstrings1(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += count(s, i, i, len(s)-1)
		sum += count(s, i, i+1, len(s)-1)
	}
	return sum
}

func count(s string, i, j, n int) int {
	sum := 0
	for i >= 0 && j <= n && s[i] == s[j] {
		sum++
		i--
		j++
	}
	return sum
}
