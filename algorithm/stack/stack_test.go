package stack

import (
	"fmt"
	"testing"
)

/*
@Time : 2020/1/18 22:50
@Author : hejun
*/

func TestArrayStack(t *testing.T) {
	stack := &ArrayStack{}
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	fmt.Println(stack.arr)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Push("1")
	stack.Push("2")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}

func TestLinkedStack(t *testing.T) {
	stack := NewLinkedListStack()
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Push("1")
	stack.Push("2")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}

func TestValidator_Check(t *testing.T) {
	v := NewValidator()
	f := v.Check("[{}()]")
	fmt.Println("check result: ", f)
	f = v.Check("[[{}()]")
	fmt.Println("check result: ", f)
}
