package _241_Different_Ways_to_Add_Parentheses

import "strconv"

func diffWaysToCompute(input string) []int {
	return genRecursive(input)
}

func genRecursive(input string) []int {
	var ret []int
	if num, err := strconv.Atoi(input); err == nil {
		ret = append(ret, num)
		return ret
	}
	for i := 1; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			left := genRecursive(input[0:i])
			right := genRecursive(input[i+1:])
			for _, l := range left {
				for _, r := range right {
					ret = append(ret, calc(l, input[i], r))
				}
			}
		}
	}
	return ret
}

func calc(l int, op byte, r int) int {
	var ret int
	switch op {
	case '+':
		ret = l + r
	case '-':
		ret = l - r
	case '*':
		ret = l * r
	}
	return ret
}
