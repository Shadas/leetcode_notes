package _189_Rotate_Array

func rotate(nums []int, k int) {
	offset := k % len(nums)
	ret := []int{}
	ret = append(nums[len(nums)-offset:], nums[0:len(nums)-offset]...)
	copy(nums, ret)
}
