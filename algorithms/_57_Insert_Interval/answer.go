package _57_Insert_Interval

import (
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	sort.Sort(X(intervals))
	var (
		ret [][]int
		idx int

		notFound bool
	)
	// 先找插入位置
	for {
		if idx >= len(intervals) {
			notFound = true
			break
		}
		cur := intervals[idx]
		if newInterval[0] < cur[0] {
			break
		} else {
			if newInterval[0] > cur[1] {
				idx++
				continue
			} else {
				break
			}
		}
	}
	// 未找到位置，直接追在最后
	if notFound {
		ret = append(intervals, newInterval)
		return ret
	}
	//fmt.Printf("insert pos=%d\n", idx)
	// 找到位置后面需要判断的第一个为idx
	// 先把前面的部分合并进来
	x := intervals[0:idx]
	ret = append(ret, x...)
	var (
		tailIdx  int
		foundPos bool
	)
	for idx < len(intervals) {
		cur := intervals[idx]
		if newInterval[0] < cur[0] && newInterval[1] < cur[0] {
			foundPos = true
			ret = append(ret, newInterval)
			tailIdx = idx
			break
		}
		// 要合并
		if cur[0] < newInterval[0] {
			newInterval[0] = cur[0]
		}
		if cur[1] > newInterval[1] {
			newInterval[1] = cur[1]
		}
		idx++
	}
	//fmt.Printf("insert tailIdx=%d, foundPos=%v\n", tailIdx, foundPos)
	// 合并尾巴
	if !foundPos {
		ret = append(ret, newInterval)
	} else {
		ret = append(ret, intervals[tailIdx:]...)
	}
	return ret
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
