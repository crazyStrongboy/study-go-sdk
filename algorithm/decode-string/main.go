package main

import "fmt"

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
}

func decodeString1(s string) string {
	result := ""
	i := 0
	mul := 0
	for i < len(s) {
		if s[i] >= 'a' && s[i] <= 'z' {
			result += string(s[i])
		} else if s[i] >= '0' && s[i] <= '9' {
			if mul == 0 {
				mul = int(s[i] - '0')
			} else {
				mul = mul*10 + int(s[i]-'0')
			}
		} else if s[i] == '[' {
			end := findEnd(s, i)
			tmp := decodeString(s[i+1 : end])
			xx := ""
			for mul >= 1 {
				xx += tmp
				mul--
			}
			result += xx
			i = end
			mul = 0
		}
		i++
	}
	return result
}

func findEnd(s string, index int) int {
	var stack []int
	stack = append(stack, 1)
	end := index + 1
	for index < len(s) && len(stack) > 0 {
		if s[end] == '[' {
			stack = append(stack, 1)
		} else if s[end] == ']' {
			stack = stack[:len(stack)-1]
		}
		end++
	}
	return end - 1
}
