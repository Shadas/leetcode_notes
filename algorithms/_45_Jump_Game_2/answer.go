package _45_Jump_Game_2

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	var (
		distance = len(nums) - 1
		step     = 0
		max      = 0
		i, idx   = 0, 0
	)
	for i < len(nums) {
		if i+nums[i] >= distance {
			step++
			return step
		}
		max = 0
		idx = i + 1
		for j := i + 1; j-i <= nums[i]; j++ {
			if max < nums[j]+j-i {
				max = nums[j] + j - i
				idx = j
			}
		}
		i = idx
		step++
	}
	return step
}
