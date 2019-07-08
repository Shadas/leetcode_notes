package _448_Find_All_Numbers_Disappeared_in_an_Array

func findDisappearedNumbers(nums []int) []int {
	l := len(nums)
	ret := []int{}
	for i := 0; i < l; i++ {
		if nums[i] <= l && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
	for i, n := range nums {
		if i+1 != n {
			ret = append(ret, i+1)
		}
	}
	return ret
}
