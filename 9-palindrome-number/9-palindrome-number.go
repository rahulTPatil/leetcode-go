package leetcode

import "strconv"

func isPalindrome(x int) bool {
	l := len(strconv.Itoa(x))

	for i := range l / 2 {
		if strconv.Itoa(x)[i] != strconv.Itoa(x)[l-i-1] {
			return false
		}
	}
	return true
}
