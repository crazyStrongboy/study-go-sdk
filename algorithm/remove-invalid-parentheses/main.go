package main

import "fmt"

func main() {
	fmt.Println(removeInvalidParentheses("(a)())()"))
}

// 广度搜索
func removeInvalidParentheses(s string) []string {
	var result []string
	set := make(map[string]struct{})
	set[s] = struct{}{}
	for len(set) > 0 {
		for k, _ := range set {
			if isValid(k) {
				result = append(result, k)
			}
		}
		if len(result) > 0 {
			return result
		}
		xx := make(map[string]struct{})
		for temp, _ := range set {
			for i, _ := range temp {
				if i > 0 && temp[i] == temp[i-1] {
					continue
				}
				if temp[i] == '(' || temp[i] == ')' {
					next := temp[:i] + temp[i+1:]
					xx[next] = struct{}{}
				}
			}
		}
		set = xx
	}
	return result
}

func isValid(s string) bool {
	cnt := 0
	for _, v := range s {
		if v == '(' {
			cnt++
		} else if v == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	if cnt > 0 {
		return false
	}
	return true
}
