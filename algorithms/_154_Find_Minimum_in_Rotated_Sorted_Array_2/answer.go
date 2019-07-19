package _154_Find_Minimum_in_Rotated_Sorted_Array_2

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var curr int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= curr {
			curr = nums[i]
		} else {
			return nums[i]
		}
	}
	return nums[0]
}
