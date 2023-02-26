package main

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	t := &T{m: map[int]int{}}
	for k, v := range inorder {
		t.m[v] = k
	}
	return t.rebuild(preorder, inorder, 0, 0, len(inorder)-1)
}

func (t *T) rebuild(preorder []int, inorder []int, root, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	rootV := preorder[root]
	index := t.m[rootV]
	node := &TreeNode{
		Val: rootV,
	}
	// 1  ,3,4,5,6,7,8,  9,10, 11,12
	// root =7  l = 8 r = 11
	node.Left = t.rebuild(preorder, inorder, root+1, l, index-1)
	node.Right = t.rebuild(preorder, inorder, root+(index-l)+1, index+1, r)
	return node
}

type T struct {
	m map[int]int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
