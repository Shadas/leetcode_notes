package _409_Longest_Palindrome

func longestPalindrome(s string) int {
	return longestPalindromeMapCount(s)
}

func longestPalindromeMapCount(s string) int {
	m := make(map[int32]int)
	for _, x := range s {
		m[x] += 1
	}
	total := 0
	hasSingle := false
	for _, v := range m {
		if v%2 == 0 {
			total += v
		} else {
			hasSingle = true
			total += v / 2 * 2
		}
	}
	if hasSingle {
		total += 1
	}

	return total
}
