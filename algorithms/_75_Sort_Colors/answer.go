package _75_Sort_Colors

func sortColors(nums []int) {
	fastSort1(nums)
}

func fastSort1(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid, i := nums[0], 1
	head, tail := 0, len(nums)-1

	for head < tail {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	nums[head] = mid
	fastSort1(nums[0:head])
	fastSort1(nums[head+1:])
}
