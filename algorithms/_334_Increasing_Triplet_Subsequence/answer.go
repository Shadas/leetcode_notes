package _334_Increasing_Triplet_Subsequence

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	//return increasingTripletDP(nums)
	return increasingTripletSimple(nums)
}

func increasingTripletSimple(nums []int) bool {
	n1, n2 := math.MaxInt, math.MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] <= n1 {
			n1 = nums[i]
		} else if nums[i] <= n2 {
			n2 = nums[i]
		} else {
			return true
		}
	}
	return false
}

// time limit exceeded with [1,1,1,1,1.....1,3,7]
func increasingTripletDP(nums []int) bool {
	dp := make([]int, len(nums))
	for idx := range dp {
		dp[idx] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			if dp[i] == 3 {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
