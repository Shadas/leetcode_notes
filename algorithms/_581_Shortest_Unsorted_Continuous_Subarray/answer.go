package _581_Shortest_Unsorted_Continuous_Subarray

import "math"

func findUnsortedSubarray(nums []int) int {
	return findUnsortedSubarrayDoubleDirection(nums)
}

func findUnsortedSubarrayDoubleDirection(nums []int) int {
	min, max := math.MinInt, math.MaxInt
	start, end := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= min {
			min = nums[i]
		} else {
			end = i
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= max {
			max = nums[i]
		} else {
			start = i
		}
	}
	if end == start {
		return 0
	}
	return end - start + 1
}
