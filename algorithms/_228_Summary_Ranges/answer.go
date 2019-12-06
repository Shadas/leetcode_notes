package _228_Summary_Ranges

import "fmt"

func summaryRanges(nums []int) []string {
	var (
		ret = []string{}
		tmp = []int{}
		idx int
	)
	for idx < len(nums) {
		if len(tmp) == 0 {
			tmp = append(tmp, nums[idx])
			idx++
		} else {
			last := tmp[len(tmp)-1]
			curr := nums[idx]
			if curr-last == 1 {
				tmp = append(tmp, curr)
				idx++
			} else {
				ret = append(ret, slice2Str(tmp))
				tmp = tmp[0:0]
			}
		}
	}
	if len(tmp) != 0 {
		ret = append(ret, slice2Str(tmp))
	}
	return ret
}

func slice2Str(tmp []int) string {
	if len(tmp) == 0 {
		return ""
	}
	if len(tmp) == 1 {
		return fmt.Sprint(tmp[0])
	}
	return fmt.Sprintf("%d->%d", tmp[0], tmp[len(tmp)-1])
}
