package _55_Jump_Game

import "math"

func canJump(nums []int) bool {
	// return canJumpDP(nums)
	return canJumpGreedy(nums)
}

func canJumpDP(nums []int) bool {
	length := len(nums)
	// 0为不可达，1为可达
	dp := make([]int, length)
	dp[len(dp)-1] = 1
	for i := length - 2; i >= 0; i-- {
		furthestJump := int(math.Min(float64(i+nums[i]), float64(length-1)))
		for j := i + 1; j <= furthestJump; j++ {
			if dp[j] == 1 {
				dp[i] = 1
				break
			}
		}
	}
	return dp[0] == 1
}

func canJumpGreedy(nums []int) bool {
	length := len(nums)
	lastPos := length - 1
	for i := length - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}
