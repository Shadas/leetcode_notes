package _392_Is_Subsequence

func isSubsequence(s string, t string) bool {
	return isSubsequenceByte(s, t)
}

func isSubsequenceByte(s string, t string) bool {
	if s == "" {
		return true
	}
	for si, ti := 0, 0; ti < len(t); ti++ {
		if t[ti] == s[si] {
			si++
		}
		if si == len(s) {
			return true
		}

	}
	return false

}
