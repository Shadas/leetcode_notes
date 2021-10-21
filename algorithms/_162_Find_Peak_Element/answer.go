package _162_Find_Peak_Element

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	// 边界
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}
