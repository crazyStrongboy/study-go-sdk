package sort

import (
	"fmt"
	"testing"
)

/*
@Time : 2020/1/19
@Author : hejun
*/

func TestInsertAscSort(t *testing.T) {
	arr := []int{6, 5, 1, 9}
	InsertAscSort(arr)
	fmt.Println(arr)
}

func TestBubbleAscSort(t *testing.T) {
	arr := []int{2, 5, 1, 9}
	BubbleAscSort(arr)
	fmt.Println(arr)
}

func TestQuickAscSort(t *testing.T) {
	arr := []int{8, 5, 1, 9, 3}
	QuickAscSort(arr)
	fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {
	arr := []int{0, 8, 5, 1, 9, 3}
	HeapSort(arr)
	fmt.Println(arr)
}
