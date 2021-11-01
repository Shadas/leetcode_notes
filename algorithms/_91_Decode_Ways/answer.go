package _91_Decode_Ways

func numDecodings(s string) int {
	return numDecodingsWithDP(s)
}

// DP 解法
func numDecodingsWithDP(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	var dp = make([]int, len(s)+1)
	dp[0] = 1 // 开始时只有1种
	dp[1] = 1 // 到达第一位只有1种方法
	for i := 2; i <= len(s); i++ {
		x := s[i-1]
		if x != '0' { // 如果此单字符合法，则它的个数至少为上一位的个数
			dp[i] = dp[i-1]
		}
		xx := s[i-2]
		if xx == '1' || (xx == '2' && x <= '6') { // 如果两位在 [10, 26]中，则还可以加上上上一位的个数
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
