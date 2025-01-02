package leetcode

func lengthOfLongestSubstring(s string) int {
	hashMap := make(map[int]int)

	for i, v := range s {
		if _, ok := hashMap[int(v)]; ok {
			return len(hashMap)
		}
		hashMap[i] = int(v)
	}
	return len(hashMap)
}
