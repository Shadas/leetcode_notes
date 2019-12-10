package _989_Add_to_Array_Form_of_Integer

func addToArrayForm(A []int, K int) []int {
	return addToArrayFormWithByte(A, K)
}

func addToArrayFormWithByte(A []int, K int) []int {
	var (
		idx = len(A) - 1
		tk  = K
		add int
		sum []int
	)
	for idx >= 0 && tk >= 0 {
		a1 := A[idx]
		b1 := tk % 10
		s := a1 + b1 + add
		sum = append([]int{s % 10}, sum...)
		add = s / 10
		idx--
		tk = tk / 10
	}
	if tk == 0 {
		for add != 0 && idx >= 0 {
			s := A[idx] + add
			sum = append([]int{s % 10}, sum...)
			add = s / 10
			idx--
		}
		if idx < 0 {
			if add > 0 {
				sum = append([]int{1}, sum...)
			}
		} else {
			sum = append(A[0:idx+1], sum...)
		}
	} else {
		for add != 0 && tk >= 0 {
			s := tk%10 + add
			sum = append([]int{s % 10}, sum...)
			add = s / 10
			tk = tk / 10
		}
		if tk == 0 {
			if add > 0 {
				sum = append([]int{1}, sum...)
			}
		} else {
			for tk > 0 {
				sum = append([]int{tk % 10}, sum...)
				tk = tk / 10
			}
		}
	}
	return sum
}
