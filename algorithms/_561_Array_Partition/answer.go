package _561_Array_Partition

import "sort"

func arrayPairSum(nums []int) int {
	return arrayPairSumWithSort(nums)
}

func arrayPairSumWithSort(nums []int) int {
	sort.Ints(nums)
	x := 0
	for i := 0; i < len(nums); i += 2 {
		x += nums[i]
	}
	return x
}
