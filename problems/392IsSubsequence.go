package leetcode

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	l := 0
	r := 0
	for i := 0; i < len(t); i++ {
		if l >= len(s) {
			break
		}
		if s[l] == t[r] {
			l++
		}
		r++
	}
	if l != len(s) {
		return false
	}
	return true
}
