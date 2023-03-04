package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("abab", "ab"))
}

func findAnagrams(s string, p string) []int {
	var result []int
	for i := 0; i < len(s)-len(p)+1; i++ {
		if isAnagrams(s[i:i+len(p)], p) {
			result = append(result, i)
		}
	}
	return result
}

func isAnagrams(s1, s2 string) bool {
	size := len(s1)
	a := make([]int, 26)
	for i := 0; i < size; i++ {
		a[s1[i]-'a']++
	}
	for i := 0; i < size; i++ {
		a[s2[i]-'a']--
	}
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}
