package queue

/*
@Time : 2020/1/19
@Author : hejun
*/
type CycleQueue struct {
	arr        []interface{}
	size       int // 环形队列的长度
	count      int // 环形队列中元素的数量
	head, tail int // 头指针和尾指针
}

func NewCycleQueue(size int) *CycleQueue {
	return &CycleQueue{
		size: size,
		arr:  make([]interface{}, size, size),
	}
}

func (q *CycleQueue) Enqueue(value interface{}) bool {
	if q.count == q.size {
		return false
	}
	q.arr[q.tail] = value
	q.tail++
	if q.tail == q.size {
		q.tail = 0
	}
	q.count++
	return true
}

func (q *CycleQueue) Dequeue() (value interface{}) {
	if q.count > 0 {
		value = q.arr[q.head]
		q.head++
		if q.head == q.size {
			q.head = 0
		}
		q.count--
	}

	return
}
