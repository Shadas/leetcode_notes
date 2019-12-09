package _415_Add_Strings

import "fmt"

func addStrings(num1 string, num2 string) string {
	return addStringsWithByte(num1, num2)
}

func addStringsWithByte(num1 string, num2 string) string {
	var (
		idx1, idx2 = len(num1) - 1, len(num2) - 1
		add        int
		sum        string
	)
	for idx1 >= 0 && idx2 >= 0 {
		a1, b1 := num1[idx1], num2[idx2]
		s := int(a1-'0') + int(b1-'0') + add
		sum = fmt.Sprint(s%10) + sum
		add = s / 10
		idx1--
		idx2--
	}
	if idx2 < 0 {
		idx2 = idx1
		num2 = num1
	}
	for idx2 >= 0 {
		s := int(num2[idx2]-'0') + add
		sum = fmt.Sprint(s%10) + sum
		add = s / 10
		idx2--
	}
	if add != 0 {
		sum = "1" + sum
	}
	return sum
}
