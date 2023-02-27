package main

func main() {
	c := Constructor(3)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)
	c.Get(4)
	c.Get(3)
	c.Get(2)
	c.Get(1)
	c.Put(5, 5)
	c.Get(1)
	c.Get(2)
	c.Get(3)
	c.Get(4)
	c.Get(5)
}

type LRUCache struct {
	m        map[int]*v
	capacity int
	head     *linked
	tail     *linked
}

type linked struct {
	prev *linked
	next *linked
	key  int
}

type v struct {
	linked *linked
	value  int
}

func Constructor(capacity int) LRUCache {
	h := &linked{}
	t := &linked{prev: h}
	h.next = t
	return LRUCache{
		m:        make(map[int]*v),
		capacity: capacity,
		head:     h,
		tail:     t,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		v.linked.prev.next, v.linked.next.prev = v.linked.next, v.linked.prev
		t := &linked{key: key, prev: this.tail.prev, next: this.tail}
		this.tail.prev.next = t
		this.tail.prev = t
		v.linked = t
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.linked.prev.next, v.linked.next.prev = v.linked.next, v.linked.prev
		t := &linked{key: key, prev: this.tail.prev, next: this.tail}
		this.tail.prev.next = t
		this.tail.prev = t
		v.value = value
		v.linked = t
		return
	}
	if len(this.m) == this.capacity {
		delete(this.m, this.head.next.key)
		this.head.next = this.head.next.next
		this.head.next.prev = this.head
	}
	t := &linked{key: key, prev: this.tail.prev, next: this.tail}
	this.tail.prev.next = t
	this.tail.prev = t
	this.m[key] = &v{
		linked: t,
		value:  value,
	}
}
