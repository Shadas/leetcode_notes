package fastsort

import (
	"testing"

	. "leetcode_notes/utils/array"
)

func TestArrayFastSort1(t *testing.T) {
	if ret := ArrayFastSort1([]int{2, 3, 4, 1, 0, 7}); !IsIntArrayEqual(ret, []int{0, 1, 2, 3, 4, 7}) {
		t.Errorf("Test1 error with return: %v", ret)
	}
}

func TestArrayFastSort2(t *testing.T) {
	if ret := ArrayFastSort2([]int{2, 3, 4, 1, 0, 7}); !IsIntArrayEqual(ret, []int{0, 1, 2, 3, 4, 7}) {
		t.Errorf("Test2 error with return: %v", ret)
	}
}
