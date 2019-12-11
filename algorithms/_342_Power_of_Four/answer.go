package _342_Power_of_Four

func isPowerOfFour(num int) bool {
	return isPowerOfFourLoop(num)
}

func isPowerOfFourLoop(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%4 != 0 {
			return false
		}
		n = n / 4
	}
	return true
}
