package leetcode

import "strings"

func lengthOfLastWord(s string) int {
	sd := strings.Fields(s)
	return len(sd[len(sd)-1])
}
