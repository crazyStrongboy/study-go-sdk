package string_match

/*
@Time : 2020/2/10
@Author : hejun
*/
const Size = 256

// BM算法进行匹配
func bm_match(origin, sub string) bool {
	bc := generateBC(sub)
	i := 0
	for i <= len(origin)-len(sub) {
		j := 0
		for j = len(sub) - 1; j >= 0; j-- {
			if origin[i+j] != sub[j] {
				break
			}
		}
		if j < 0 {
			return true
		}
		i += j - bc[origin[i+j]]
	}
	return false
}

func generateBC(sub string) [Size]int {
	bc := [Size]int{}
	for i := 0; i < Size; i++ {
		bc[i] = -1
	}
	for i := 0; i < len(sub); i++ {
		bc[sub[i]] = i
	}
	return bc
}
