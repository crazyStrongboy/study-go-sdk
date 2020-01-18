package stack

import "fmt"

/*
@Time : 2020/1/18 23:01
@Author : hejun
*/
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

type Validator struct {
	stack *ArrayStack
}

func NewValidator() *Validator {
	return &Validator{stack: &ArrayStack{}}
}

func IsOpen(value Bracket) bool {
	return value == OpenRoundBracket || value == OpenBracket || value == OpenBrace
}

func (v *Validator) Check(value string) bool {
	v.stack.Reset()
	vs := []rune(value)
	for _, r := range vs {
		word := string(r)
		if IsOpen(Bracket(word)) {
			v.stack.Push(Bracket(word))
		} else {
			if p := v.stack.Pop(); p == nil || validMap[p.(Bracket)] != Bracket(word) {
				return false
			}
		}
	}
	fmt.Println(v.stack.arr)
	// 扫描完还有未匹配的OpenBracket，返回false
	if len(v.stack.arr) > 0 {
		return false
	}
	return true
}
