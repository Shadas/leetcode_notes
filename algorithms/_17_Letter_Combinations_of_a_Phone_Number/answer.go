package _17_Letter_Combinations_of_a_Phone_Number

func letterCombinations(digits string) []string {
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
	//return letterCombinationsLoop(dict, digits)
	return letterCombinationsDfs(dict, digits)
}

func letterCombinationsDfs(dict map[string][]string, digits string) []string {
	if digits == "" {
		return []string{}
	}
	ret := []string{}
	dfs(dict, &ret, 0, "", digits)
	return ret
}

func dfs(dict map[string][]string, ret *[]string, idx int, tmp, digits string) {
	if idx == len(digits) {
		*ret = append(*ret, tmp)
		return
	}
	for _, v := range dict[string(digits[idx])] {
		tmp += v
		dfs(dict, ret, idx+1, tmp, digits)
		tmp = tmp[0 : len(tmp)-1]
	}
}

func letterCombinationsLoop(dict map[string][]string, digits string) []string {
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
