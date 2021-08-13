package _17_Letter_Combinations_of_a_Phone_Number

func letterCombinations(digits string) []string {

	return letterCombinationsLoop(digits)
}

func letterCombinationsLoop(digits string) []string {
	dict := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	ret := []string{}
	for _, v := range digits {
		ret = foo(dict, ret, string(v))
	}
	return ret
}

func foo(dict map[string][]string, ret []string, n string) []string {
	if len(ret) == 0 {
		return dict[n]
	}
	re := []string{}
	for _, r := range ret {
		for _, v := range dict[n] {
			re = append(re, r+v)
		}
	}
	return re
}
