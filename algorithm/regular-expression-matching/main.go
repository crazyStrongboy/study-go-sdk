package main

import "fmt"

func main() {
	fmt.Println(isMatch("ba", ".*."))
}

func isMatch(s string, p string) bool {
	i := 0
	j := 0
	for j < len(p) && i < len(s) {
		if j+1 < len(p) && p[j+1] == '*' {
			if isMatch(s[i:], p[j+2:]) { // 进行回退
				return true
			}
		}
		if p[j] == '*' {
			tmp := p[j-1]
			if tmp != s[i] && tmp != '.' {
				j++
			} else {
				i++
			}
			continue
		}
		if p[j] == '.' {
			i++
			j++
			continue
		}
		if s[i] != p[j] {
			return false
		}
		i++
		j++
	}
	if i < len(s) {
		// 后续都是相同元素，并且p只剩最后一个'*'
		//if j == len(p)-1 && p[j] == '*' {
		//	unique := s[i-1]
		//	for i < len(s) {
		//		if s[i] != unique {
		//			return false
		//		}
		//		i++
		//	}
		//	return true
		//}
		// 后续只有一个元素，并且p只剩一个'.'
		if j == len(p)-1 && p[j] == '.' && i == len(s)-1 {
			return true
		}
		return false
	}
	if j < len(p) {
		if !isBack(p[j:]) {
			return false
		}
		if p[len(p)-1] == '*' {
			return true
		}
		if p[len(p)-1] == s[len(s)-1] && p[j] == '*' {
			return true
		}
		return false
	}
	return true
}

func isBack(s string) bool {
	cnt := 0
	for i := range s {
		if s[i] != '*' {
			cnt++
			if cnt > 1 {
				return false
			}
		} else {
			cnt--
		}
	}
	return cnt <= 1
}
