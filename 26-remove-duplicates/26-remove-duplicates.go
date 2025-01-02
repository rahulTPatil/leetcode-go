package leetcode

func removeDuplicates(nums []int) int {
	hashMap := make(map[int]int, len(nums))
	count := 0
	hashMap[0] = nums[0]
	for i, v := range nums {
		if _, ok := hashMap[v]; ok {
			count++
			continue
		}
		hashMap[i] = v
	}
	return count
}
