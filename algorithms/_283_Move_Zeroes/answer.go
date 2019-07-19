package _283_Move_Zeroes

func moveZeroes(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			continue
		}
		// 如果发现0了,向后找第一个不是0的交换，
		j := i + 1
		for j <= len(nums)-1 && nums[j] == 0 {
			j++
		}
		if j >= len(nums) { // 后面全是0
			return
		}
		// 找到不是0的，交换
		nums[i], nums[j] = nums[j], nums[i]
	}
}
