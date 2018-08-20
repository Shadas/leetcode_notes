package _80_Remove_Duplicates_from_Sorted_Array_2

func removeDuplicates(nums []int) int {
	var (
		index      int
		lastValue  int
		repeatTime int
	)
	if len(nums) == 0 {
		return 0
	}
	lastValue = nums[index]
	index++
	repeatTime = 1

	for {
		if len(nums) <= index {
			break
		}
		if nums[index] == lastValue {
			repeatTime++
			if repeatTime > 2 {
				if index == (len(nums) - 1) {
					nums = nums[0:index]
				} else {
					nums = append(nums[0:index], nums[index+1:]...)
				}
			} else {
				index++
			}

		} else {
			lastValue = nums[index]
			repeatTime = 1
			index++
		}
	}

	return len(nums)
}
