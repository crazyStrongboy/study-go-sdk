package main

import "fmt"

type Trie struct {
	m     map[string]bool
	array [26]*elem
}

type elem struct {
	val [26]*elem
}

func Constructor() Trie {
	return Trie{
		m:     make(map[string]bool),
		array: [26]*elem{},
	}
}

func (this *Trie) Insert(word string) {
	this.m[word] = true
	var e *elem
	first := word[0] - 'a'
	e = this.array[first]
	if e == nil {
		e = &elem{val: [26]*elem{}}
		this.array[first] = e
	}
	for i := 1; i < len(word); i++ {
		index := word[i] - 'a'
		//fmt.Println(index)
		if e.val[index] == nil {
			e.val[index] = &elem{val: [26]*elem{}}
		}
		e = e.val[index]
	}
}

func (this *Trie) Search(word string) bool {
	return this.m[word]
}

func (this *Trie) StartsWith(prefix string) bool {
	if this.array[prefix[0]-'a'] == nil {
		return false
	}
	e := this.array[prefix[0]-'a']
	for i := 1; i < len(prefix); i++ {
		e = e.val[prefix[i]-'a']
		if e == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	c := Constructor()
	c.Insert("apple")
	fmt.Println(c.Search("apple"))
	fmt.Println(c.Search("app"))
	fmt.Println(c.StartsWith("app"))
	//fmt.Println(c.array)
}
