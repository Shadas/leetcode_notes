package _647_Palindromic_Substrings

func countSubstrings(s string) int {
	return countSubstringsWithDP(s)
}

// dp[i][j] 代表s[i,j]是否为回文子串
// 如果s[i,j]为true，且 s[i-1] == s[j+1]，s[i-1,j+1]也为true
// 相当于计算dp矩阵里，包含对角线的右上半边的true数量
func countSubstringsWithDP(s string) int {
	var (
		count int
		dp    [][]bool
	)
	dp = make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true // 单个字符肯定为回文，初始化对角线，也是初始化长度为1的
		count++
	}
	// 第二轮，初始化长度为2的
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			dp[i-1][i] = true
			count++
		}
	}
	// 开始遍历dp数组
	for i := 2; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[j-i] == s[j] && dp[j-i+1][j-1] {
				dp[j-i][j] = true
				count++
			}
		}
	}
	// 观察dp特征用
	//for _, line := range dp {
	//	fmt.Println(line)
	//}
	return count
}
