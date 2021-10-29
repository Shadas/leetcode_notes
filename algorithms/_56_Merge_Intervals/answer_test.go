package _56_Merge_Intervals

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	var (
		intervals, ret [][]int
	)
	intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	ret = merge(intervals)
	fmt.Println(ret)

	intervals = [][]int{{1, 4}, {4, 5}}
	ret = merge(intervals)
	fmt.Println(ret)

	intervals = [][]int{{1, 4}, {1, 5}}
	ret = merge(intervals)
	fmt.Println(ret)
}
