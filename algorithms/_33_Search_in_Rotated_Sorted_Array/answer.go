package _33_Search_in_Rotated_Sorted_Array

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 && nums[0] == target {
		return 0
	}
	var (
		left  = 0
		right = len(nums) - 1
		mid   = 0
	)
	for left <= right {
		mid = (left + right) / 2
		// 如果碰上了就碰上了
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[left] { // 左边单调递增，右边不一定
			if nums[mid] < target { // 肯定不在左边的单调区间，左边界调整，右边界不定，所以这轮右边界不变
				left = mid + 1
			} else { // 比mid小，可能在左边，也可能在右边rotate的部分，需要根据左边下界判断
				if nums[left] > target { // 如果单调区间下界也比target大，则目标值一定落在右边部分
					right = right - 1
					left = mid + 1
				} else { // 如果单调区间下界比target小，说明就在这个单调区间中
					right = mid - 1
					left = left + 1
				}
			}
		} else { // 左边不单调递增，则右边一定单调递增
			if nums[mid] > target { // 说明右边单调递增部分都比目标值大，所以右界调整，左界+1继续处理
				right = mid - 1
				left = left + 1
			} else { // 说明在mid右边的单调区间，或者 左边rotate部分，需要根据右边上界判断处理
				if nums[right] > target { // 如果右边界大于目标值，说明就在右边的单调区间中
					right = right - 1
					left = mid + 1
				} else { // 否则，在左边比mid小的单调区间中
					left = left + 1
					right = mid - 1
				}
			}
		}

	}
	return -1
}
