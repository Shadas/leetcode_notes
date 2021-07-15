package _132_Palindrome_Partitioning_2

import (
	"math"
)

func minCut(s string) int {
	n := len(s)
	judge := make([][]bool, n)
	for idx := range judge {
		judge[idx] = make([]bool, n)
	}
	dp := make([]int, n) // dp[i]为s中第i位到第n-1位的子字符串中，最小分割次数
	for i := n - 1; i >= 0; i-- {
		dp[i] = math.MaxInt32
		for j := i; j < n; j++ {
			if s[j] == s[i] && (j-i <= 1 || judge[i+1][j-1]) {
				judge[i][j] = true
				if j+1 < n {
					dp[i] = int(math.Min(float64(dp[i]), 1+float64(dp[j+1])))
				} else {
					dp[i] = 0
				}
			}
		}
	}
	return dp[0]
}
