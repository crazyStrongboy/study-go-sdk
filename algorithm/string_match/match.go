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

var letter_map = map[uint8]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
	'i': 8,
	'j': 9,
	'k': 10,
	'l': 11,
	'm': 12,
	'n': 13,
	'o': 14,
	'p': 15,
	'q': 16,
	'r': 17,
	's': 18,
	't': 19,
	'u': 20,
	'v': 21,
	'w': 22,
	'x': 23,
	'y': 24,
	'z': 25,
}

// RK算法进行匹配，a-z对应0-26
func rk_match(origin, sub string) bool {
	// 计算出子串对应的十进制值
	subSum := 0
	for i := 0; i < len(sub); i++ {
		subSum += letter_map[sub[i]]
	}
	tempSum := 0
	count := 0
	tempIndex := 0
	head := letter_map[origin[tempIndex]]
	for i := 0; i < len(origin); i++ {
		if count < len(sub) {
			tempSum += letter_map[origin[i]]
			count++
		}
		if count == len(sub) && tempSum == subSum {
			// next check，进行完全匹配，房子计算出来的值冲突
			tempSub := origin[tempIndex : tempIndex+len(sub)]
			if completeMatch(tempSub, sub) {
				return true
			}
		}
		if count == len(sub) {
			count--
			tempSum -= head
			tempIndex++
			head = letter_map[origin[tempIndex]]
		}
	}
	return false
}

// 进行完全匹配
func completeMatch(a, b string) bool {
	l := len(a)
	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
