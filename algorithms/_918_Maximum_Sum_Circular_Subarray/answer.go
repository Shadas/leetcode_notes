package _918_Maximum_Sum_Circular_Subarray

func maxSubarraySumCircular(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		sum    = nums[0]
		curMax = nums[0]
		curMin = nums[0]
		maxSum = nums[0]
		minSum = nums[0]
		n      = len(nums)
	)
	for i := 1; i < n; i++ {
		sum += nums[i]
		curMax = max(curMax+nums[i], nums[i])
		curMin = min(curMin+nums[i], nums[i])
		maxSum = max(maxSum, curMax)
		minSum = min(minSum, curMin)
	}
	if minSum == sum {
		return maxSum
	}
	return max(maxSum, sum-minSum)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
