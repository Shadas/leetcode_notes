package _204_Count_Primes

func countPrimes(n int) int {
	return countPrimesMultiplication(n)
	// return countPrimesDivision(n)
}

func countPrimesMultiplication(n int) int {
	var (
		table = make([]int, n)
		ret   int
	)
	for i := 2; i < n; i++ {
		if table[i] != 0 {
			continue
		}
		ret++
		table[i] = 1
		for j := 1; j*i < n; j++ {
			table[j*i] = -1
		}
	}
	return ret
}

// 会超时
func countPrimesDivision(n int) int {
	var (
		table = make([]int, n)
		ret   int
	)
	for i := 2; i < n; i++ {
		if table[i] != 0 {
			continue
		}
		ret++
		table[i] = 1
		for j := i + 1; j < n; j++ {
			if table[j] == 0 && j%i == 0 {
				table[j] = -1
			}
		}
	}
	return ret
}
