package _4_Median_of_Two_Sorted_Arrays

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	if ret := findMedianSortedArrays([]int{1}, []int{1}); ret != 1 {
		t.Error("err test1 want 1 with", ret)
	}
	if ret := findMedianSortedArrays([]int{1}, []int{2}); ret != 1.5 {
		t.Error("err test2 want 1.5 with", ret)
	}
	if ret := findMedianSortedArrays([]int{1, 2}, []int{3, 4}); ret != 2.5 {
		t.Error("err test3 want 2.5 with", ret)
	}
}
