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

func isMatchDP(s, p string) bool {
	return false
}
