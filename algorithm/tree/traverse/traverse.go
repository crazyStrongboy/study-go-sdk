package traverse

import (
	"fmt"
	"github.com/crazyStrongboy/study-go-sdk/algorithm/tree"
)

/*
@Time : 2020/2/4
@Author : hejun
*/

// 前序，先打印任意节点
func preTraverse(root *tree.Node) {
	if root == nil {
		return
	}
	fmt.Println(root)
	preTraverse(root.Left)
	preTraverse(root.Right)
}

// 中序，中间打印任意节点
func midTraverse(root *tree.Node) {
	if root == nil {
		return
	}
	midTraverse(root.Left)
	fmt.Println(root)
	midTraverse(root.Right)
}

// 后序，最后打印任意节点
func postTraverse(root *tree.Node) {
	if root == nil {
		return
	}
	postTraverse(root.Left)
	postTraverse(root.Right)
	fmt.Println(root)
}
