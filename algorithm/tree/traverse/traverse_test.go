package traverse

import (
	"github.com/crazyStrongboy/study-go-sdk/algorithm/tree"
	"testing"
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
