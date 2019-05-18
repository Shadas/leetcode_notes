package _202_Happy_Number

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	record := []int{}
	return isHappyA(n, record)
}

func isHappyA(n int, record []int) bool {
	m := happy(n)
	if m == 1 {
		return true
	}
	for _, r := range record {
		if m == r {
			return false
		}
	}
	record = append(record, m)
	return isHappyA(m, record)
}

func happy(n int) int {
	sum := 0
	for {
		tmp := n % 10
		sum += tmp * tmp
		n = n / 10
		if n == 0 {
			break
		}
	}
	return sum
}
