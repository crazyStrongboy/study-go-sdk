package main

import "github.com/crazyStrongboy/study-go-sdk/algorithm/linked"

/*
@Time : 2020/1/17
@Author : hejun
*/

type lru struct {
	total int
	count int
	l     *linked.ListNode
}

// 新增一个节点
func (lru *lru) add(value string) {
	if !lru.remove(value) && lru.count >= lru.total {
		// 移除尾部的节点
		head := lru.l
		for {
			next := head.Next
			if next == nil {
				lru.l = nil
				break
			}
			if next.Next == nil {
				head.Next = nil
				break
			}
			head = next
		}
	}
	head := linked.ListNode{Value: value}
	head.Next = lru.l
	lru.l = &head
	lru.count++
}

// 移除对应的元素
func (lru *lru) remove(value string) bool {
	if lru.l == nil {
		return true
	}
	// 判断头节点是否是要移除的节点
	if lru.l.Value == value {
		lru.l = lru.l.Next
		lru.count--
		return true
	}
	var prev *linked.ListNode = lru.l
	var temp *linked.ListNode = lru.l.Next
	for {
		if temp == nil {
			return false
		}
		if temp.Value == value {
			break
		}
		prev = temp
		temp = temp.Next
	}
	prev.Next = temp.Next
	lru.count--
	return true
}

func main() {
	lru := lru{total: 2}
	lru.add("a")
	lru.add("b")
	lru.add("c")
	lru.add("d")
	lru.l.String()
}
