package main

import "fmt"

func main() {
	fmt.Println(minWindow("aa", "aa"))
}

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	needNum := 0
	for i := range t {
		need[t[i]]++
	}
	left := 0
	right := 0
	subS := s + "_"
	for left <= len(s)-len(t) && right < len(s) {
		if _, ok := need[s[right]]; !ok {
			right++
			continue
		}
		need[s[right]]--
		if need[s[right]] == 0 {
			needNum++ // 该字符刚好收集好，记录字符搜集的次数
		}
		if needNum == len(need) {
			for needNum == len(need) {
				if len(subS) > right-left {
					subS = s[left : right+1]
				}
				if _, ok := need[s[left]]; ok {
					need[s[left]]++
					if need[s[left]] > 0 {
						needNum--
					}
				}
				left++
				continue
			}
		}
		right++
	}
	if subS == s+"_" {
		return ""
	}
	return subS
}
