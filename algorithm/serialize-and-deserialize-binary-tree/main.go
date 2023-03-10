package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	c := &Codec{}
	xx := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	ss := c.serialize(xx)
	fmt.Println(ss)
	r := c.deserialize(ss)
	fmt.Println(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	queue []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "_#"
	}
	result := fmt.Sprintf("_%d", root.Val)
	result += this.serialize(root.Left)
	result += this.serialize(root.Right)
	return result
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "_#" {
		return nil
	}
	data = data[1:]
	arr := strings.Split(data, "_")
	this.queue = arr
	return this.helper()
}

func (this *Codec) helper() *TreeNode {
	if len(this.queue) == 0 {
		return nil
	}
	v := this.queue[0]
	this.queue = this.queue[1:]
	if v == "#" {
		return nil
	}
	root := &TreeNode{Val: cInt(v)}
	root.Left = this.helper()
	root.Right = this.helper()
	return root
}

func cInt(s string) int {
	xx, _ := strconv.ParseInt(s, 10, 64)
	return int(xx)
}
