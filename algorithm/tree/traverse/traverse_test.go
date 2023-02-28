package traverse

import (
	"fmt"
	"testing"

	"github.com/crazyStrongboy/study-go-sdk/algorithm/tree"
)

func Test_preTraverse(t *testing.T) {
	root := tree.NewNode(2)
	root.Left = tree.NewNode(1)
	root.Right = tree.NewNode(3)
	root.Right.Right = tree.NewNode(4)
	preTraverse(root)
}

func Test_midTraverse(t *testing.T) {
	root := tree.NewNode(2)
	root.Left = tree.NewNode(1)
	root.Right = tree.NewNode(3)
	root.Right.Right = tree.NewNode(4)
	midTraverse(root)
}

func Test_postTraverse(t *testing.T) {
	root := tree.NewNode(2)
	root.Left = tree.NewNode(1)
	root.Right = tree.NewNode(3)
	root.Right.Right = tree.NewNode(4)
	postTraverse(root)
}

func Test_iterMid(t *testing.T) {
	root := tree.NewNode(2)
	root.Left = tree.NewNode(1)
	root.Right = tree.NewNode(3)
	root.Right.Right = tree.NewNode(4)
	midTraverse(root)
	fmt.Println(result)
	fmt.Println(iterMid(root))
}

func Test_iterPre(t *testing.T) {
	root := tree.NewNode(2)
	root.Left = tree.NewNode(1)
	root.Right = tree.NewNode(3)
	root.Right.Right = tree.NewNode(4)
	preTraverse(root)
	fmt.Println(result)
	fmt.Println(iterPre(root))
}

func Test_iterPost(t *testing.T) {
	root := tree.NewNode(2)
	root.Left = tree.NewNode(1)
	root.Right = tree.NewNode(3)
	root.Right.Right = tree.NewNode(4)
	postTraverse(root)
	fmt.Println(result)
	fmt.Println(iterPost(root))
}
