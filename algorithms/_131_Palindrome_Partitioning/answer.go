package _131_Palindrome_Partitioning

func partition(s string) [][]string {
	var ret [][]string
	backTrack(s, 0, []string{}, &ret)
	return ret
}

func isPalindromic(str string) bool {
	for i := 0; i <= len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func backTrack(str string, idx int, tmp []string, ret *[][]string) {
	if idx == len(str) {
		var x = make([]string, len(tmp))
		for idx, v := range tmp {
			x[idx] = v
		}
		*ret = append(*ret, x)
		return
	}
	for i := idx; i < len(str); i++ {
		sub := str[idx : i+1]
		if isPalindromic(sub) {
			tmp = append(tmp, sub)
			backTrack(str, i+1, tmp, ret)
			tmp = tmp[0 : len(tmp)-1]
		}
	}
}
