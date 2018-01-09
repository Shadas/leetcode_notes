package _53_Maximum_Subarray

import (
	"math"
)

func maxSubArray(nums []int) int {
	return maxSubArrayDP1(nums)
}

func maxSubArrayDP1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < n; i++ {
		var x int
		if dp[i-1] > 0 {
			x = dp[i-1]
		} else {
			x = 0
		}
		dp[i] = nums[i] + x
		max = int(math.Max(float64(max), float64(dp[i])))
	}
	return max
}
