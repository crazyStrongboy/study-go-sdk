package main

import "fmt"

type Trie struct {
	m     map[string]bool
	array [26]*elem
}

type elem struct {
	val  [26]int
	next *elem
}

func (e *elem) String() string {
	return fmt.Sprintf("val: %v-next: %v\n", e.val, e.next)
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
		e = &elem{
			val: [26]int{},
		}
		this.array[first] = e
	}
	for i := 1; i < len(word); i++ {
		index := word[i] - 'a'
		fmt.Println(index)
		e.val[index] = 1
		if e.next == nil {
			e.next = &elem{
				val: [26]int{},
			}
		}
		e = e.next
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
		if e != nil && e.val[prefix[i]-'a'] != 1 {
			return false
		}
		e = e.next
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
