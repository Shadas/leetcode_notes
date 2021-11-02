package _209_Minimum_Size_Subarray_Sum

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	return minSubArrayLenWithSlidingWindow(target, nums)
}

func minSubArrayLenWithSlidingWindow(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		p   int
		q   int = 1
		sum int = nums[p]
		ret int = math.MaxInt32
	)
	for p < len(nums) || q < len(nums) {
		//fmt.Printf("p=%d, q=%d, sum=%d\n", p, q, sum)
		if target <= sum {
			tmp := q - p
			if tmp < ret {
				ret = tmp
			}
			if target == sum {
				if q >= len(nums) {
					break
				} else {
					sum += nums[q]
					q += 1
				}
			} else { // 大于target的情形
				if p >= len(nums) {
					break
				} else {
					sum -= nums[p]
					p += 1
				}
			}

		} else if target > sum {
			if q >= len(nums) {
				break
			}
			sum += nums[q]
			q += 1
		}
	}
	if ret == math.MaxInt32 {
		ret = 0
	}
	return ret
}
