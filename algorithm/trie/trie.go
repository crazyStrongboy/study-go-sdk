package trie

/*
@Time : 2020/2/13
@Author : hejun
*/

type TrieNode struct {
	data     int32
	children [26]*TrieNode
	isEnd    bool
}

func node(data int32) *TrieNode {
	return &TrieNode{data: data}
}

var root = node('/')

func insert(s string) {
	p := root
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if p.children[index] == nil {
			p.children[index] = node(int32(s[i]))
		}
		p = p.children[index]
	}
	p.isEnd = true
}

func find(s string) bool {
	p := root
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if p.children[index] == nil {
			return false
		}
		p = p.children[index]
	}
	if p.isEnd {
		return true
	}
	return false
}
