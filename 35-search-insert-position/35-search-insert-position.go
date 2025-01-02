package leetcode

import "sort"

func searchInsert(nums []int, target int) int {
	nums = append(nums, target)
	sort.Ints(nums)
	idx := 0

	for i, v := range nums {
		if v == target {
			idx = i
			return idx
		}
	}
	return idx
}
