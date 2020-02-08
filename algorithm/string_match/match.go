package string_match

/*
@Time : 2020/2/8
@Author : hejun
*/

// BF算法进行匹配
func bf_match(origin, sub string) bool {
outer:
	for i := 0; i < len(origin)-len(sub); i++ {
		for j := 0; j < len(sub); j++ {
			if sub[j] != origin[i+j] {
				continue outer
			}
		}
		return true
	}
	return false
}
