package traverse

import (
	"fmt"

	"github.com/crazyStrongboy/study-go-sdk/algorithm/tree"
)

/*
@Time : 2020/2/4
@Author : hejun
*/
var result []interface{}

// 前序，先打印任意节点
func preTraverse(root *tree.Node) {
	if root == nil {
		return
	}
	result = append(result, root.Elem)
	preTraverse(root.Left)
	preTraverse(root.Right)
}

// 中序，中间打印任意节点
func midTraverse(root *tree.Node) {
	if root == nil {
		return
	}
	midTraverse(root.Left)
	result = append(result, root.Elem)
	midTraverse(root.Right)
}

// 后序，最后打印任意节点
func postTraverse(root *tree.Node) {
	if root == nil {
		return
	}
	postTraverse(root.Left)
	postTraverse(root.Right)
	result = append(result, root.Elem)
}

func iterMid(root *tree.Node) []interface{} {
	var (
		stack []*tree.Node
		r     []interface{}
	)
	stack = append(stack, root)
	for len(stack) != 0 {
		elem := stack[len(stack)-1]
		if elem != nil {
			stack = stack[:len(stack)-1]
			if elem.Right != nil {
				stack = append(stack, elem.Right)
			}
			stack = append(stack, elem)
			stack = append(stack, nil)
			if elem.Left != nil {
				stack = append(stack, elem.Left)
			}
			continue
		}
		elem = stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		r = append(r, elem.Elem)
	}
	return r
}

func iterPre(root *tree.Node) []interface{} {
	var (
		stack []*tree.Node
		r     []interface{}
	)
	stack = append(stack, root)
	for len(stack) != 0 {
		elem := stack[len(stack)-1]
		if elem != nil {
			stack = stack[:len(stack)-1]
			if elem.Right != nil {
				stack = append(stack, elem.Right)
			}
			if elem.Left != nil {
				stack = append(stack, elem.Left)
			}
			stack = append(stack, elem)
			stack = append(stack, nil)
			continue
		}
		elem = stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		r = append(r, elem.Elem)
	}
	return r
}

func iterPost(root *tree.Node) []interface{} {
	var (
		stack []*tree.Node
		r     []interface{}
	)
	stack = append(stack, root)
	for len(stack) != 0 {
		elem := stack[len(stack)-1]
		if elem != nil {
			stack = stack[:len(stack)-1]

			if elem.Left != nil {
				stack = append(stack, elem.Left)
			}
			if elem.Right != nil {
				stack = append(stack, elem.Right)
			}
			stack = append(stack, elem)
			stack = append(stack, nil)

			continue
		}
		elem = stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		r = append(r, elem.Elem)
	}
	fmt.Println(r)
	i := 0
	j := len(r) - 1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return r
}
