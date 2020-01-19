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
		arr[j+1] = value
	}
}

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
