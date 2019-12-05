package _263_Ugly_Number

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	rules := []int{2, 3, 5}
	for _, r := range rules {
		for num%r == 0 {
			num = num / r
		}
	}
	return num == 1
}
