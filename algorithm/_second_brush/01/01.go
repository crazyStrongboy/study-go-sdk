package main

import (
	"fmt"
	"sort"
)

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

func longestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		left := i
		right := i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > len(result) {
				result = s[left : right+1]
			}
			left--
			right++
		}
		left = i
		right = i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > len(result) {
				result = s[left : right+1]
			}
			left--
			right++
		}

	}
	return result
}

func isMatch(s string, p string) bool {
	var dp [][]bool
	for i := 0; i <= len(s); i++ {
		dp = append(dp, make([]bool, len(p)+1))
	}
	dp[0][0] = true // 都是空，天然true
	for j := 2; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2] // 空串初始化dp数组
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					zero := dp[i][j-2]
					one := dp[i-1][j-2]
					many := dp[i-1][j]
					dp[i][j] = zero || one || many
				} else {
					dp[i][j] = dp[i][j-2]
				}

			}
		}
	}
	return dp[len(s)][len(p)]
}

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	result := 0
	for i < j {
		result = max(result, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}

	}
	return result
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	t := &T{}
	t.backtrack(digits, 0, []byte{})
	return t.result
}

var letters = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

type T struct {
	result []string
}

func (t *T) backtrack(digits string, start int, track []byte) {
	if start >= len(digits) {
		//fmt.Println(string(track))
		t.result = append(t.result, string(track))
		return
	}

	letter := letters[digits[start]]
	for j := range letter {
		track = append(track, letter[j])
		t.backtrack(digits, start+1, track)
		track = track[:len(track)-1]
	}

}

// threeSum1 三数之和第一版
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := 0 - (nums[i] + nums[j])
			if v, ok := m[target]; ok {
				if v > j {
					result = append(result, []int{nums[i], nums[j], target})
				}

			}
		}
	}
	return result
}
