package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	left := (m + n + 1) / 2
	right := (m + n + 2) / 2
	return float64(findTopK(nums1, 0, nums2, 0, left)+
		findTopK(nums1, 0, nums2, 0, right)) / 2
}

func findTopK(nums1 []int, start1 int, nums2 []int, start2 int, k int) int {
	if start1 >= len(nums1) {
		return nums2[start2+k-1]
	}
	if start2 >= len(nums2) {
		return nums1[start1+k-1]
	}
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}
	mid1 := start1 + k/2 - 1
	if mid1 >= len(nums1) {
		mid1 = len(nums1) - 1
	}
	mid2 := start2 + k/2 - 1
	if mid2 >= len(nums2) {
		mid2 = len(nums2) - 1
	}
	if nums1[mid1] < nums2[mid2] {
		return findTopK(nums1, start1+k/2, nums2, start2, k-min(k/2, len(nums1)-start1))
	} else {
		return findTopK(nums1, start1, nums2, start2+k/2, k-min(k/2, len(nums2)-start2))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
