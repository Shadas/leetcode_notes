package _13_Roman_to_Integer

func romanToInt(s string) int {
	return romanToIntScan(s)
}

func romanToIntScan(s string) int {
	if len(s) == 0 {
		return 0
	}
	var (
		sum int
		pre int
	)
	pre = c2i(s[0])
	for idx := 1; idx < len(s); idx++ {
		num := c2i(s[idx])
		if pre < num {
			sum -= pre
		} else {
			sum += pre
		}
		pre = num
	}
	sum += pre

	return sum
}

func c2i(x uint8) int {
	var i int
	switch x {
	case 'I':
		i = 1
	case 'V':
		i = 5
	case 'X':
		i = 10
	case 'L':
		i = 50
	case 'C':
		i = 100
	case 'D':
		i = 500
	case 'M':
		i = 1000
	}
	return i
}
