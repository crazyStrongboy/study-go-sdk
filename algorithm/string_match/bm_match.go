package string_match

/*
@Time : 2020/2/10
@Author : hejun
*/
const Size = 256

// BM算法进行匹配
func bm_match(origin, sub string) bool {
	bc := generateBC(sub)
	suffix, prefix := generateGS(sub)
	i := 0
	for i <= len(origin)-len(sub) {
		j := 0
		for j = len(sub) - 1; j >= 0; j-- {
			if origin[i+j] != sub[j] { // 坏字符在i+j位置处
				break
			}
		}
		if j < 0 {
			return true
		}
		x := j - bc[origin[i+j]] // 判断坏字符是否在匹配串中
		y := 0
		if j < len(sub)-1 {
			y = moveByGs(j, len(sub), suffix, prefix)
		}
		i += max(x, y)
	}
	return false
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func moveByGs(j int, m int, suffix []int, prefix []bool) int {
	k := m - 1 - j // 好后缀长度
	if suffix[k] != -1 {
		return j - suffix[k] + 1 // 有好后缀
	}
	// 这里r需要等于j+2。例如后缀为abc，如果没有匹配好后缀，则需要从bc开始匹配
	for r := j + 2; r < m-1; r++ {
		if prefix[m-r] { // 与好后缀匹配的前缀
			return r
		}
	}
	return m
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

func generateGS(sub string) (suffix []int, prefix []bool) {
	l := len(sub)
	for i := 0; i < l; i++ {
		suffix = append(suffix, -1)
		prefix = append(prefix, false)
	}
	for i := 0; i < len(sub)-1; i++ {
		j := i
		k := 0 // 后缀子串长度
		for j >= 0 && sub[j] == sub[l-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}
		if j == -1 {
			prefix[i] = true
		}
	}
	return
}
