package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	fmt.Println(m)
	for i := 0; i < len(nums); i++ {
		fmt.Println(m[target-nums[i]])
		if v, ok := m[target-nums[i]]; ok && i != v {
			return []int{i, v}
		}
	}
	return []int{}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	pre := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + pre
		v := sum % 10
		pre = sum / 10
		head.Next = &ListNode{Val: v}
		head = head.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + pre
		v := sum % 10
		pre = sum / 10
		head.Next = &ListNode{Val: v}
		head = head.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + pre
		v := sum % 10
		pre = sum / 10
		head.Next = &ListNode{Val: v}
		head = head.Next
		l2 = l2.Next
	}

	if pre != 0 {
		head.Next = &ListNode{Val: pre}
	}

	return dummy.Next
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	m := make(map[byte]int)
	j := 0
	for i := 0; i < len(s); i++ {
		for j < len(s) && m[s[j]] <= 0 {
			m[s[j]] += 1
			result = max(result, j-i+1)
			j++
		}
		m[s[i]]--
	}
	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	left := (l1 + l2 + 1) / 2
	right := (l1 + l2 + 2) / 2
	return float64(findMedianValue(nums1, 0, nums2, 0, left)+findMedianValue(nums1, 0, nums2, 0, right)) / 2
}

func findMedianValue(nums1 []int, start1 int, nums2 []int, start2 int, k int) int {
	if start1 >= len(nums1) {
		return nums2[start2+k-1]
	}
	if start2 >= len(nums2) {
		return nums1[start1+k-1]
	}
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}
	mid1 := min(start1+k/2-1, len(nums1)-1)
	mid2 := min(start2+k/2-1, len(nums2)-1)
	//fmt.Println(mid1,mid2)
	if nums1[mid1] < nums2[mid2] {
		return findMedianValue(nums1, start1+k/2, nums2, start2, max(k-k/2, k-(len(nums1)-start1)))
	} else {
		return findMedianValue(nums1, start1, nums2, start2+k/2, max(k-k/2, k-(len(nums2)-start2)))
	}
}
