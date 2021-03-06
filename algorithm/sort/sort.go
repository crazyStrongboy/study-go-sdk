package sort

/*
@Time : 2020/1/19
@Author : hejun
*/
// 插入算法进行升序排序
func InsertAscSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ {
		value := arr[i] //要排序的元素
		j := i - 1      // 前面均为有序队列
		for ; j >= 0; j-- {
			if arr[j] > value {
				// 往后移动一位
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		// 空出来的那一个位置，把value填进去
		arr[j+1] = value
	}
}

// 冒泡排序
func BubbleAscSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr); i++ {
		flag := false // 提前退出冒泡排序的标志
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
				flag = true // 表示没有进行数据交换
			}
		}
		if !flag {
			break // 没有数据交换，提前退出
		}
	}
}

// 快速排序
func QuickAscSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickAscSort(arr, 0, len(arr)-1)
}

// 进行快速排序
func quickAscSort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	mid := partition(arr, start, end)
	quickAscSort(arr, start, mid-1)
	quickAscSort(arr, mid+1, end)
}

// 获取分区索引
func partition(arr []int, start int, end int) int {
	// 取end节点为分区节点，小于它的排在它左边，大于它的排在它右边，原地排序
	v := arr[end]
	j := start
	for i := start; i < end; i++ {
		if arr[i] < v {
			swap(arr, i, j)
			j++
		}
	}
	swap(arr, j, end)
	return j
}

// 交换数组中的两个元素
func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

// 堆排序
func HeapSort(arr []int) {
	// 建堆
	buildHeap(arr)
	// 取出1位置元素放到最后一个，对前面的n-1个元素重新建堆，直到只有一个元素为止
	n := len(arr) - 1
	for n > 1 {
		swap(arr, 1, n)
		n--
		heapify(arr, 1, n)
	}
}

// 建堆
func buildHeap(arr []int) {
	n := len(arr)
	for i := n / 2; i >= 1; i-- {
		heapify(arr, i, len(arr)-1)
	}
}

// 堆化
func heapify(arr []int, i, n int) {
	for {
		max := i
		if i*2 <= n && arr[max] < arr[i*2] {
			max = i * 2
		}
		if i*2+1 <= n && arr[max] < arr[i*2+1] {
			max = i*2 + 1
		}
		if max == i {
			break
		}
		swap(arr, i, max)
		i = max
	}
}
