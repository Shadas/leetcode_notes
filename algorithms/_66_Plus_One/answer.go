package _66_Plus_One

func plusOne(digits []int) []int {
	l := len(digits)
	var isAdd bool
	var ret []int
	for i := l - 1; i >= 0; i-- {
		v := digits[i]
		if i == l-1 || isAdd {
			v = v + 1
			if v == 10 {
				isAdd = true
				v = 0
			} else {
				isAdd = false
			}
		}
		ret = append([]int{v}, ret...)
	}
	if isAdd {
		ret = append([]int{1}, ret...)
	}
	return ret
}
