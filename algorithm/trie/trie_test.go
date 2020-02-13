package trie

import (
	"fmt"
	"testing"
)

func Test_insert(t *testing.T) {
	insert("hello")
	insert("hi")
	insert("how")
	insert("see")
	fmt.Println(root)
}

func Test_find(t *testing.T) {
	insert("hello")
	insert("hi")
	insert("how")
	insert("see")

	fmt.Println(find("how"))
	fmt.Println(find("hii"))
}
