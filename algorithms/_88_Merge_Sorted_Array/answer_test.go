package _88_Merge_Sorted_Array

import (
	. "leetcode_notes/utils/array"
	"testing"
)

func TestMerge(t *testing.T) {
	var nums1, nums2 []int
	nums1 = []int{0}
	nums2 = []int{1}
	if merge(nums1, 0, nums2, 1); !IsIntArrayEqual(nums1, []int{1}) {
		t.Error("error, get", nums1)
	}
	nums1 = []int{1, 0}
	nums2 = []int{2}
	if merge(nums1, 1, nums2, 1); !IsIntArrayEqual(nums1, []int{1, 2}) {
		t.Error("error, get", nums1)
	}
}
