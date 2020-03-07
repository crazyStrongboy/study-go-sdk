package main

import "fmt"

/*
@Time : 2020/3/7
@Author : hejun
*/

/**
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 num1 成为一个有序数组。



说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。


示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

*/

func main() {
	num1 := []int{2, 0}
	merge(num1, 1, []int{1}, 1)
	fmt.Println(num1)
}

// 从后往前遍历
func merge(nums1 []int, m int, nums2 []int, n int) {
	len1 := m - 1
	len2 := n - 1
	l := m + n - 1
	for len1 >= 0 && len2 >= 0 {
		if nums2[len2] > nums1[len1] {
			nums1[l] = nums2[len2]
			len2--
		} else {
			nums1[l] = nums1[len1]
			len1--
		}
		l--
	}
	for i := 0; i <= len2; i++ {
		nums1[i] = nums2[i]
	}
}
