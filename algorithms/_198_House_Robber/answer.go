package _198_House_Robber

import "math"

func rob(nums []int) int {
	// return solution1(len(nums)-1, nums)
	// return solution2(nums)
	return solution3(nums)
}

// 纯递归，time limit exceeded
func solution1(idx int, nums []int) int {
	if idx < 0 {
		return 0
	}
	max := int(math.Max(float64(nums[idx]+solution1(idx-2, nums)), float64(solution1(idx-1, nums))))
	return max
}

// 带结果储存的递归
func solution2(nums []int) int {
	var ret []int
	for i := 0; i < len(nums); i++ {
		ret = append(ret, -1)
	}
	return solution2Recursion(len(nums)-1, nums, ret)
}

func solution2Recursion(idx int, nums []int, ret []int) int {
	if idx < 0 {
		return 0
	}
	if idx < len(ret) && ret[idx] != -1 {
		return ret[idx]
	}
	max := int(math.Max(float64(nums[idx]+solution2Recursion(idx-2, nums, ret)), float64(solution2Recursion(idx-1, nums, ret))))
	ret[idx] = max
	return max
}

// 自底向上，DP  dp[i] = max(money[i]+dp[i-2], dp[i-1])
func solution3(nums []int) int {
	var ret []int
	for i := 0; i < len(nums); i++ {
		ret = append(ret, -1)
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	ret[0] = nums[0]
	ret[1] = int(math.Max(float64(ret[0]), float64(nums[1])))
	for i := 2; i < len(nums); i++ {
		ret[i] = int(math.Max(float64(nums[i]+ret[i-2]), float64(ret[i-1])))
	}
	return ret[len(nums)-1]
}
