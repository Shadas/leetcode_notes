package _44_Wildcard_Matching

func isMatch(s string, p string) bool {
	return isMatchScan(s, p)
}

func isMatchScan(s string, p string) bool {
	var (
		i, j    int          // s,p的idx
		lp          = len(p) // 表达式长度
		starPos int = -1     // 记录*的idx，-1代表没遇到
		iStar   int          // 匹配*过程中的i的位置
	)
	for i < len(s) { // s 遍历完之前
		if j < lp && (p[j] == '?' || p[j] == s[i]) { // 如果正常匹配，可以继续检查
			i += 1
			j += 1
		} else if j < lp && p[j] == '*' { // 如果匹配到了星星，记录星星位置，从星星下一个开始匹配，遍历i的后面一段
			starPos = j
			j += 1
			iStar = i // 之后从s这个位置开始匹配*的过程
		} else if starPos != -1 { // 在star的匹配过程中
			// 如果进第二次进入*的状态，说明前面的*已经匹配成功了，所以关心最新一个*就行
			// 注意这里的if是有顺序的，进入这个状态后：
			//// 如果成功匹配了*后面第一位，会在上面向后继续走
			////// 遇到新的*处理就行，结束就结束了
			//// 如果走着走不下去了，说明当时第一个*后匹配可能匹配早了：
			////// j重置到*后第一位准备匹配;
			////// iStar向后+1，从脱离*匹配的地方向后继续走*匹配的过程
			j = starPos + 1 // 在star匹配过程中，匹配字符串的对比量永远在*后第一位
			iStar += 1
			i = iStar
		} else {
			return false
		}
	}
	// 处理j不到尾部的情形
	for j < lp && p[j] == '*' {
		j++
	}
	return j == lp // 如果能对齐，则匹配成功
}

// 状态转移方程
// if p[j-1] = '*'; f(i,j) = f(i, j-1) or f(i-1, j)
//// 为*的时候，就考虑从当前位置向后匹配0次还是1次
////// 如果是0次，则，f(i-1, j)
////// 如果是1次，则，f(i, j-1), 在这个情况下，后续如果再匹配，则在当前计算中，再向后找
// if p[j-1] !-= '*'; f(i,j) = f(i-1, j-1) and (s[i-1] == p[j-1] or p[j-1] == '?')
//// 不为*的时候，前面的匹配成功，且当前能够单个字符对应
func isMatchDP(s, p string) bool {
	var dp [][]bool // [s][p]
	// 初始化dp数组
	dp = make([][]bool, len(s)+1)
	for idx := 0; idx <= len(s); idx++ {
		dp[idx] = make([]bool, len(p)+1)
	}
	dp[0][0] = true // 初始长度0，0的时候是匹配的
	//
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				if i > 0 {
					dp[i][j] = dp[i][j-1] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			} else {
				if i > 0 {
					dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '?')
				} else {
					dp[i][j] = false
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
