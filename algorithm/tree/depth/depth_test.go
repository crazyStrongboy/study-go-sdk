package depth

import (
	"fmt"
	"testing"

	"github.com/crazyStrongboy/study-go-sdk/algorithm/tree"
)

func Test_minDepth(t *testing.T) {
	root := tree.NewNode(2)
	root.Right = tree.NewNode(3)
	root.Right.Left = tree.NewNode(1)
	root.Right.Right = tree.NewNode(4)
	root.Right.Right.Right = tree.NewNode(5)

	fmt.Println(minDepth(root))
}
