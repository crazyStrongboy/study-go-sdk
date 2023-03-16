package main

func subsets(nums []int) [][]int {
	s := &subSet{used: make(map[int]bool)}
	s.backtrack(nums, 0, []int{})
	return s.result
}

type subSet struct {
	result [][]int
	used   map[int]bool
}

func (s *subSet) backtrack(nums []int, start int, track []int) {
	t := make([]int, len(track))
	copy(t, track)
	s.result = append(s.result, t)
	for i := start; i < len(nums); i++ {
		// if s.used[nums[i]]{
		//     continue
		// }
		track = append(track, nums[i])
		// s.used[nums[i]] = true
		s.backtrack(nums, i+1, track)
		track = track[:len(track)-1]
		// s.used[nums[i]] = false
	}
}
