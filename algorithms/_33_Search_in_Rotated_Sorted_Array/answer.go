package _33_Search_in_Rotated_Sorted_Array

import "fmt"

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
		fmt.Println(left, mid, right)
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[left] {
			if nums[mid] < target {
				left = mid + 1
			} else {
				if nums[left] > target {
					right = right - 1
					left = mid + 1
				} else {
					right = mid - 1
					left = left + 1

				}
			}
		} else {
			if nums[mid] > target {
				right = mid - 1
				left = left + 1
			} else {
				if nums[right] > target {
					right = right - 1
					left = mid + 1
				} else {
					left = left + 1
					right = mid - 1
				}
			}
		}

	}
	return -1
}
