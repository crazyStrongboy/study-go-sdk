package main

func main() {

}

func leastInterval(tasks []byte, n int) int {
	m := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		m[tasks[i]]++
	}
	maxLen := 0
	count := 1
	for _, v := range m {
		if maxLen < v {
			maxLen = v
			count = 1
			continue
		}
		if v == maxLen {
			count++
		}
	}
	return max((maxLen-1)*(n+1)+count, len(tasks))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
