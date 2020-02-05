package operate

import (
	"fmt"
)

/*
@Time : 2020/2/5
@Author : hejun
*/

var root *node

func find(data int) *node {
	if root == nil {
		return nil
	}
	curr := root
	for curr != nil {
		if curr.elem == data {
			return curr
		} else if data > curr.elem {
			curr = curr.right
		} else {
			curr = curr.left
		}
	}
	return nil
}

func insert(data int) {
	if root == nil {
		root = newNode(data)
		return
	}
	curr := root

	for {
		if data > curr.elem {
			if curr.right == nil {
				curr.right = newNode(data)
				return
			}
			curr = curr.right
		} else {
			if curr.left == nil {
				curr.left = newNode(data)
				return
			}
			curr = curr.left
		}
	}

}

func delete(data int) {
	curr := root     // 要删除的节点
	var parent *node // 要删除的节点的父节点
	for curr != nil {
		if curr.elem == data {
			break
		} else if data > curr.elem {
			parent = curr
			curr = curr.right
		} else {
			parent = curr
			curr = curr.left
		}
	}
	// 在当前树未查询到节点
	if curr == nil {
		return
	}
	// 左右子节点均存在，将右节点的最终左子节点与该节点交换，然后移除最终左子节点
	if curr.left != nil && curr.right != nil {
		p := curr.right
		pp := parent
		for p.left != nil {
			pp = p
			p = p.left
		}
		curr.elem = p.elem
		curr = p
		parent = pp
	}
	var child *node
	if curr.left != nil {
		child = curr.left
	} else if curr.right != nil {
		child = curr.right
	} else {
		child = nil
	}

	if parent == nil {
		root = child // 移除的是root节点
	} else if parent.left == curr {
		parent.left = child
	} else {
		parent.right = child
	}
}

type node struct {
	elem        int
	left, right *node
}

func newNode(elem int) *node {
	return &node{elem: elem}
}

// 中序，中间打印任意节点
func midTraverse(root *node) {
	if root == nil {
		return
	}
	midTraverse(root.left)
	fmt.Println(root)
	midTraverse(root.right)
}

func (node *node) String() string {
	return fmt.Sprintf("%d", node.elem)
}
