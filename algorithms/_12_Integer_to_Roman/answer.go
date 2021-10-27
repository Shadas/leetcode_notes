package _12_Integer_to_Roman

func intToRoman(num int) string {
	return intToRomanGreedy(num)
}

func intToRomanGreedy(num int) string {
	var (
		str     string
		values  = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	)
	for idx, value := range values {
		for value <= num {
			num -= value
			str += symbols[idx]
		}
	}
	return str
}
