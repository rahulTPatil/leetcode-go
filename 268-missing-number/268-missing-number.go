package leetcode

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	for i := range len(nums) {
		if nums[i] != i {
			return i
		}
	}
	return l
}
