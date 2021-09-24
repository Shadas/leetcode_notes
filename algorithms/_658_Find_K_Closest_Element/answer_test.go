package _658_Find_K_Closest_Element

import (
	"testing"

	"github.com/shadas/leetcode_notes/utils/array"
)

func TestFindClosestElements(t *testing.T) {
	var (
		arr, ret []int
		k, x     int
	)
	arr = []int{1, 2, 3, 4, 5}
	k, x = 4, 3
	ret = findClosestElements(arr, k, x)
	if !array.IsIntArrayEqual(ret, []int{1, 2, 3, 4}) {
		t.Errorf("wrong ret is %v", ret)
	}
	arr = []int{1, 1, 1, 10, 10, 10}
	k, x = 1, 9
	ret = findClosestElements(arr, k, x)
	if !array.IsIntArrayEqual(ret, []int{10}) {
		t.Errorf("wrong ret is %v", ret)
	}
	arr = []int{1, 1, 2, 2, 2, 2, 2, 3, 3}
	k, x = 3, 3
	ret = findClosestElements(arr, k, x)
	if !array.IsIntArrayEqual(ret, []int{2, 3, 3}) {
		t.Errorf("wrong ret is %v", ret)
	}
}
