package _56_Merge_Intervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	return mergeAfterSort(intervals)
}

func mergeAfterSort(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}
	sort.Sort(X(intervals))
	var (
		ret [][]int
		pre []int
	)
	pre = intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if needMerge(pre, cur) {
			pre[1] = cur[1]
		} else if needKeep(pre, cur) {
			continue
		} else {
			ret = append(ret, pre)
			pre = cur
		}
	}
	ret = append(ret, pre)
	return ret
}

func needKeep(pre, cur []int) bool {
	if cur[1] <= pre[1] && cur[0] >= pre[0] {
		return true
	}
	return false
}

func needMerge(pre, cur []int) bool {
	if pre[0] == cur[0] || pre[1] == cur[1] {
		return true
	}
	if pre[1] <= cur[1] && pre[1] >= cur[0] {
		return true
	}
	return false
}

type X [][]int

func (x X) Len() int {
	return len(x)
}

func (x X) Less(i, j int) bool {
	if x[i][0] < x[j][0] {
		return true
	} else if x[i][0] > x[j][0] {
		return false
	} else {
		if x[i][1] <= x[j][1] {
			return true
		} else {
			return false
		}
	}
}

func (x X) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
