package _258_Add_Digits

func addDigits(num int) int {
	return addDightsSmart(num)
}

func addDightsSmart(num int) int {
	return 1 + (num-1)%9
}
