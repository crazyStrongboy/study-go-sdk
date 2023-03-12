package main

import "fmt"

func main() {
	xx := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println(xx)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	result := merge(nums1, nums2)
	if len(result) == 0 {
		return 0
	}
	size := len(result)
	if len(result)%2 == 0 {
		return float64(result[size/2]+result[size/2-1]) / 2
	}
	return float64(result[len(result)/2])
}

func merge(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	var result []int
	i := 0
	j := 0
	for {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
			if i >= len(nums1) {
				break
			}
		} else {
			result = append(result, nums2[j])
			j++
			if j >= len(nums2) {
				break
			}
		}
	}
	if i < len(nums1) {
		result = append(result, nums1[i:]...)
	}
	if j < len(nums2) {
		result = append(result, nums2[j:]...)
	}
	return result
}
