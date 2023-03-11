package main

import "strings"

func decodeString(s string) string {
	ans := &strings.Builder{}
	var multistack []int
	var ansstack []*strings.Builder

	multi := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			multi = multi*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			multistack = append(multistack, multi)
			ansstack = append(ansstack, ans)
			ans = &strings.Builder{}
			multi = 0
		} else if s[i] >= 'a' && s[i] <= 'z' {
			ans.WriteByte(s[i])
		} else {
			tmpans := ansstack[len(ansstack)-1]
			ansstack = ansstack[:len(ansstack)-1]
			num := multistack[len(multistack)-1]
			multistack = multistack[:len(multistack)-1]
			for num >= 1 {
				tmpans.WriteString(ans.String())
				num--
			}
			ans = tmpans
		}
	}
	return ans.String()
}
