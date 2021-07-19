package _792_Number_of_Matching_Subsequences

func numMatchingSubseq(s string, words []string) int {
	dict := make(map[string][]int)
	for idx, b := range s {
		dict[string(b)] = append(dict[string(b)], idx)
	}
	ret := 0
	for _, word := range words {
		if isSubSeq(word, dict) {
			ret += 1
		}
	}
	return ret
}

func isSubSeq(word string, dict map[string][]int) bool {
	prev := -1
	for i := 0; i < len(word); i++ {
		s := string(word[i])
		idxs := dict[s]
		if len(idxs) == 0 { // 没有对应字符
			return false
		}
		nh := nextHigher(prev, idxs)
		if nh == -1 {
			return false
		} else {
			prev = nh
		}
	}
	return true
}

func nextHigher(n int, list []int) int {
	for _, idx := range list {
		if idx > n {
			return idx
		}
	}
	return -1
}
