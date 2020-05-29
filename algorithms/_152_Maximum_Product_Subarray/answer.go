package _152_Maximum_Product_Subarray

func maxProduct(nums []int) int {
	return maxProductDP1(nums)
}

func maxProductDP1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	currMax, currMin := make([]int, len(nums)), make([]int, len(nums))
	currMax[0], currMin[0] = nums[0], nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		currMin[i] = MultiMin(nums[i], currMin[i-1]*nums[i], currMax[i-1]*nums[i])
		currMax[i] = MultiMax(nums[i], currMax[i-1]*nums[i], currMin[i-1]*nums[i])
		max = MultiMax(max, currMax[i])
	}
	return max
}

func MultiMin(nums ...int) int {
	if len(nums) == 0 {
		panic("empty nums")
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func MultiMax(nums ...int) int {
	if len(nums) == 0 {
		panic("empty nums")
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
