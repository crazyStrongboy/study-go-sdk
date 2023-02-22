package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"abc", "bac", "aer"}))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	var ret [][]string
	for i := 0; i < len(strs); i++ {
		s := strs[i]
		ss := SortString(s)
		m[ss] = append(m[ss], s)
	}
	for _, s := range m {
		ret = append(ret, s)
	}
	return ret
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
