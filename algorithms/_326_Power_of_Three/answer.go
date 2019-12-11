package _326_Power_of_Three

func isPowerOfThree(n int) bool {
	return isPowerOfThreeLoop(n)
}

func isPowerOfThreeLoop(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n = n / 3
	}
	return true
}

func isPowerOfThreeSmart(n int) bool {
	// too many smart ways
	return false
}
