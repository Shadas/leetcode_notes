package _5_Longest_Palindromic_Substring

var location, maxLength int

func longestPalindrome(s string) string {
	location, maxLength = 0, 0
	l := len(s)
	if l < 2 {
		return s
	}
	for i := 0; i < l-1; i++ {
		extendPalindrome(s, i, i)
		extendPalindrome(s, i, i+1)
	}
	return string(s[location : location+maxLength])
}

func extendPalindrome(s string, j, k int) {
	for {
		if j >= 0 && k < len(s) && s[j] == s[k] {
			j--
			k++
		} else {
			break
		}
	}
	if length := k - j - 1; length > maxLength {
		location = j + 1
		maxLength = length
	}
}
