package _300_Longest_Increasing_Subsequence

func lengthOfLIS(nums []int) int {
	return lengthOfLISDP(nums)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLISDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		dp      = make([]int, len(nums))
		max int = 1
	)
	for idx, _ := range dp {
		dp[idx] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
		max = maxInt(max, dp[i])
	}
	return max
}
