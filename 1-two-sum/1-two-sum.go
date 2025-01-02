package leetcode

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, v := range nums {
		comp := target - v

		if idx, ok := hashMap[comp]; ok {
			return []int{idx, i}
		}

		hashMap[v] = i
	}
	return []int{-1, -1}
}
