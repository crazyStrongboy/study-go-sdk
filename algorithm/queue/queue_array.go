package queue

import "fmt"

/*
@Time : 2020/1/19
@Author : hejun
*/

type ArrayQueue struct {
	arr        []interface{}
	size       int
	head, tail int
}

func NewArrayQueue(size int) *ArrayQueue {
	return &ArrayQueue{
		size: size,
		arr:  make([]interface{}, size*2, size*2),
	}
}

func (q *ArrayQueue) Enqueue(value interface{}) bool {
	if q.tail-q.head == q.size {
		return false
	}
	q.arr[q.tail] = value
	//q.arr = append(q.arr, value)
	q.tail++
	return true
}

func (q *ArrayQueue) Dequeue() (value interface{}) {
	if q.head == q.size {
		//q.arr = q.arr[q.head:q.tail]
		copy(q.arr, q.arr[q.head:q.tail])
		q.tail -= q.head
		q.head = 0
	}
	if q.head < q.tail {
		value = q.arr[q.head]
		q.head++
		return
	}

	return
}

func (q *ArrayQueue) String() string {
	s := "["
	h := q.head
	t := q.tail
	for h < t {
		s += fmt.Sprintf("%+v", q.arr[h])
		s += " "
		h++
	}
	s += "]"
	return s
}
