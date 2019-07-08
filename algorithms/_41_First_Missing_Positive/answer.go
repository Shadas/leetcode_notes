package _41_First_Missing_Positive

import "fmt"

func firstMissingPositive(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		fmt.Println(i)
		if nums[i] > 0 && nums[i] < length && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
	for i, n := range nums {
		if i+1 != n {
			return i + 1
		}
	}
	return length + 1
}
