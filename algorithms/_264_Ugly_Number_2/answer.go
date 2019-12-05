package _264_Ugly_Number_2

func nthUglyNumber(n int) int {
	var (
		u                = make([]int, n)
		idx2, idx3, idx5 int
		k                = 1
	)
	u[0] = 1

	for k < n {
		u[k] = min(u[idx2]*2, u[idx3]*3, u[idx5]*5)
		if u[idx2]*2 == u[k] {
			idx2++
		}
		if u[idx3]*3 == u[k] {
			idx3++
		}
		if u[idx5]*5 == u[k] {
			idx5++
		}
		k++
	}
	return u[n-1]
}

func min(a, b, c int) int {
	var min int
	if a < b {
		min = a
	} else {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}
