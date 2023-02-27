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
	expire   []int
	capacity int
}

type v struct {
	index int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:        make(map[int]*v),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		i := v.index
		for _, k := range this.expire[i+1:] {
			this.m[k].index--
		}
		this.expire = append(this.expire[0:i], this.expire[i+1:]...)
		this.expire = append(this.expire, key)
		v.index = len(this.expire) - 1
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		i := v.index
		for _, k := range this.expire[i+1:] {
			this.m[k].index--
		}
		this.expire = append(this.expire[0:i], this.expire[i+1:]...)
		this.expire = append(this.expire, key)
		v.index = len(this.expire) - 1
		v.value = value
		return
	}
	if len(this.m) == this.capacity {
		expireKey := this.expire[0]
		delete(this.m, expireKey)
		this.expire = this.expire[1:]
		for _, k := range this.expire {
			this.m[k].index--
		}
	}
	this.expire = append(this.expire, key)
	this.m[key] = &v{
		index: len(this.expire) - 1,
		value: value,
	}
}
