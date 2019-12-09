package _680_Valid_Palindrome_2

func validPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	var (
		i, j    = 0, len(s) - 1
		ti, tj  int
		errFlag int = 0
	)
	for i < j {
		if s[i] != s[j] {
			errFlag = 1
			break
		}
		i++
		j--
	}
	if errFlag == 0 {
		return true
	}
	ti, tj = i, j
	// 假如删除前面的
	i++
	for i < j {
		if s[i] != s[j] {
			errFlag = 2
			break
		}
		i++
		j--
	}
	if errFlag == 1 {
		return true
	}
	// 假设删除后面的
	i, j = ti, tj
	j--
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
