package _57_Insert_Interval

import "testing"

func TestInsert(t *testing.T) {
	var (
		intervals   [][]int
		newInterval []int
		ret         [][]int
	)
	intervals, newInterval = [][]int{{1, 3}, {6, 9}}, []int{2, 5}
	ret = insert(intervals, newInterval)
	t.Logf("ret is %v", ret) // [[1,5], [6,9]]

	intervals, newInterval = [][]int{{1, 5}}, []int{2, 7}
	ret = insert(intervals, newInterval)
	t.Logf("ret is %v", ret) // [[1,7]]

	intervals, newInterval = [][]int{{1, 5}}, []int{0, 0}
	ret = insert(intervals, newInterval)
	t.Logf("ret is %v", ret) // [[0,0], [1,5]]

	intervals, newInterval = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}
	ret = insert(intervals, newInterval)
	t.Logf("ret is %v", ret) // [[1,2],[3,10],[12,16]]
}
