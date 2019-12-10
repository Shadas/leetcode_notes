package _231_Power_of_Two

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
