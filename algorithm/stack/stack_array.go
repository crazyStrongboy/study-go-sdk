package stack

/*
@Time : 2020/1/18 22:29
@Author : hejun
*/

type ArrayStack struct {
	arr []interface{}
}

func (s *ArrayStack) Push(value interface{}) {
	s.arr = append(s.arr, value)
}

func (s *ArrayStack) Pop() (value interface{}) {
	if len(s.arr) == 0 {
		return
	}
	value = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return
}

func (s *ArrayStack) Reset() {
	s.arr = s.arr[:0]
}
