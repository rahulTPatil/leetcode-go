package leetcode

func isValid(s string) bool {
	res := 0

	for i := range s {
		switch s[i] {
		case '(':
			res += 1
		case ')':
			res -= 1
		case '[':
			res += 2
		case ']':
			res -= 2
		case '{':
			res += 3
		case '}':
			res -= 3
		}
	}
	return res == 0
}
