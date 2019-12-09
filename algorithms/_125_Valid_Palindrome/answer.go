package _125_Valid_Palindrome

import "strings"

func isPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	var (
		i, j = 0, len(s) - 1
	)
	for i <= j {
		if !isWord(s[i]) {
			i++
			continue
		}
		if !isWord(s[j]) {
			j--
			continue
		}
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isWord(b byte) bool {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return true
	}
	return false
}
