package _139_Word_Break

func wordBreak(s string, wordDict []string) bool {
	//return wordBreakBackTrack(s, wordDict)
	return wordBreakDP(s, wordDict)
}

// 回溯法，会超时
func wordBreakBackTrack(s string, wordDict []string) bool {
	return backFunc(s, 0, wordDict)
}

func backFunc(s string, idx int, wordDict []string) bool {
	if idx >= len(s) {
		return true
	}
	for i := idx; i < len(s); i++ {
		subStr := s[idx : i+1]
		if wordInDict(subStr, wordDict) {
			ret := backFunc(s, i+1, wordDict)
			if ret {
				return true
			}
		}
	}
	return false
}

// DP法  dp[j] = dp[i] && (s[i:j+1] in wordDict)
func wordBreakDP(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for j := 1; j <= len(s); j++ {
		for i := 0; i < j; i++ {
			substr := s[i:j]
			if dp[i] && wordInDict(substr, wordDict) {
				dp[j] = true
			}
		}
	}
	return dp[len(s)]
}

func wordInDict(word string, dict []string) bool {
	for _, w := range dict {
		if word == w {
			return true
		}
	}
	return false
}
