package _15_3Sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || i > 0 && nums[i] != nums[i-1] {
			var header, tail, sum = i + 1, len(nums) - 1, 0 - nums[i]
			for {
				if header >= tail {
					break
				}
				if nums[header]+nums[tail] == sum {
					ret = append(ret, []int{nums[i], nums[header], nums[tail]})
					for {
						if header >= tail || nums[header] != nums[header+1] {
							break
						}
						header++
					}
					for {
						if header >= tail || nums[tail] != nums[tail-1] {
							break
						}
						tail--
					}
					header++
					tail--
				} else if nums[header]+nums[tail] < sum {
					header++
				} else {
					tail--
				}
			}
		}
	}
	return ret
}
