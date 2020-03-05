package main

import "fmt"

/*
@Time : 2020/3/5
@Author : hejun
*/

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

*/

func main() {
	valid := isValid("()[]{}")
	fmt.Println(valid)
}

type Bracket string

func (b *Bracket) String() string {
	return string(*b)
}

const (
	OpenRoundBracket  Bracket = "("
	CloseRoundBracket Bracket = ")"
	OpenBracket       Bracket = "["
	CloseBracket      Bracket = "]"
	OpenBrace         Bracket = "{"
	CloseBrace        Bracket = "}"
)

var validMap = map[Bracket]Bracket{OpenRoundBracket: CloseRoundBracket, OpenBracket: CloseBracket, OpenBrace: CloseBrace}

func IsOpen(value Bracket) bool {
	return value == OpenRoundBracket || value == OpenBracket || value == OpenBrace
}
func isValid(s string) bool {
	if s == "" {
		return true
	}
	stack := &ArrayStack{}
	for _, r := range s {
		word := string(r)
		if IsOpen(Bracket(word)) {
			stack.Push(Bracket(word))
		} else {
			if p := stack.Pop(); p == nil || validMap[p.(Bracket)] != Bracket(word) {
				return false
			}
		}
	}
	if len(stack.arr) > 0 {
		return false
	}
	return true
}

type ArrayStack struct {
	arr []interface{}
}

func (s *ArrayStack) Push(value interface{}) {
	s.arr = append(s.arr, value)
}

func (s *ArrayStack) Pop() (value interface{}) {
	if len(s.arr) == 0 {
		return
	}
	value = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return
}
