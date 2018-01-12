package _26_Remove_Duplicates_from_Sorted_Array

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[0:i], nums[i+1:]...)
			i = -1
		}
	}
	return len(nums)
}
