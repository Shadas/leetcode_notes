package _213_House_Robber_2

import "math"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var ret1, ret2 int
	ret1 = getMaxInt(nums[:len(nums)-1])
	ret2 = getMaxInt(nums[1:])
	return int(math.Max(float64(ret1), float64(ret2)))
}

func getMaxInt(nums []int) int {
	var result []int
	for i := 0; i < len(nums); i++ {
		result = append(result, -1)
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	result[0] = nums[0]
	result[1] = int(math.Max(float64(nums[1]), float64(result[0])))
	for i := 2; i < len(nums); i++ {
		result[i] = int(math.Max(float64(result[i-1]), float64(result[i-2]+nums[i])))
	}
	return result[len(result)-1]
}
