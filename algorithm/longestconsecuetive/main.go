package main

func main() {

}

func longestConsecuetive(nums []int) int {
	m := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	result := 0
	for v := range m {
		if !m[v-1] {
			curr := 1
			temp := v
			for m[temp+1] {
				curr++
				temp++
			}
			if result < curr {
				result = curr
			}
		}
	}
	return result
}
