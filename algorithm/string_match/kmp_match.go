package string_match

/*
@Time : 2020/2/11
@Author : hejun
*/
func kmp_match(origin, sub string) bool {
	next := getNext(sub)
	j := 0
	for i := 0; i < len(origin); i++ {
		for j > 0 && origin[i] != sub[j] {
			j = next[j-1] + 1
		}
		if origin[i] == sub[j] {
			j++
		}
		if j == len(sub) {
			return true
		}
	}
	return false
}

func getNext(sub string) []int {
	l := len(sub)
	next := make([]int, l, l)
	next[0] = -1
	k := -1
	for i := 1; i < l; i++ {
		for k != -1 && sub[i] != sub[k+1] {
			k = next[k]
		}
		if sub[k+1] == sub[i] {
			k++
		}
		next[i] = k
	}
	return next
}
