package _9_Palindrome_Number

func isPalindrome(x int) bool {
	var n int
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	for {
		if n >= x {
			break
		}
		n = n*10 + x%10
		x = x / 10
	}
	return n == x || n/10 == x
}
