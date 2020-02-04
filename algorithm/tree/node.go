package tree

import "fmt"

/*
@Time : 2020/2/4
@Author : hejun
*/

type Node struct {
	Elem  interface{}
	Left  *Node
	Right *Node
}

func NewNode(elem interface{}) *Node {
	return &Node{Elem: elem}
}

func (node *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", node.Elem, node.Left, node.Right)
}
