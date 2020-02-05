package operate

import (
	"fmt"
	"testing"
)

func Test_find(t *testing.T) {
	insert(10)
	insert(2)
	insert(4)
	insert(3)
	node := find(5)
	fmt.Println(node)
	node = find(2)
	fmt.Println(node)
}

func Test_insert(t *testing.T) {
	insert(10)
	insert(2)
	insert(4)
	insert(3)
	midTraverse(root)
}

func Test_delete(t *testing.T) {
	insert(2)
	insert(10)
	insert(4)
	insert(3)
	midTraverse(root)
	fmt.Println("======================")
	delete(4)
	midTraverse(root)
}
