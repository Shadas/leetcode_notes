package _22_Generate_Parentheses

func generateParenthesis(n int) []string {
	return generateParenthesisWithMap(n)
}

func generateParenthesisWithMap(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	dp := make([]map[string]struct{}, n)
	dp[0] = map[string]struct{}{
		"()": {},
	}
	for i := 2; i <= n; i++ {
		pre := dp[i-2]
		x := genNewStrs(pre)
		dp[i-1] = x
	}
	ret := make([]string, len(dp[n-1]))
	idx := 0
	for k := range dp[n-1] {
		ret[idx] = k
		idx++
	}
	return ret
}

func genNewStrs(prem map[string]struct{}) (ret map[string]struct{}) {
	ret = make(map[string]struct{})
	for k := range prem {
		for i := range k {
			pre := k[0:i]
			last := k[i:]
			tmp := pre + "()" + last
			ret[tmp] = struct{}{}
		}
	}
	return ret
}
