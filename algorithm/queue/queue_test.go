package queue

import (
	"fmt"
	"testing"
)

/*
@Time : 2020/1/19
@Author : hejun
*/

func TestArrayQueue(t *testing.T) {
	queue := NewArrayQueue(2)
	for i := 0; i < 4; i++ {
		ret := queue.Enqueue(i)
		fmt.Printf("i： %d,ret: %t\n", i, ret)
		fmt.Println(queue.Dequeue())
	}
	queue.Enqueue(5)
	fmt.Println(queue)
	for i := 0; i < 4; i++ {
		fmt.Println(queue.Dequeue())
	}
}
