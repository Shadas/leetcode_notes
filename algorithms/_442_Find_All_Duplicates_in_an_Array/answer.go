package _442_Find_All_Duplicates_in_an_Array

func findDuplicates(nums []int) []int {
	l := len(nums)
	ret := []int{}
	for i := 0; i < l; i++ {
		if nums[i] <= l && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
	for i, n := range nums {
		if i+1 != n {
			ret = append(ret, n)
		}
	}
	return ret
}
