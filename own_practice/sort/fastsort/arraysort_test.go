package fastsort

import (
	"testing"
)

func TestArrayFastSort1(t *testing.T) {
	if ret := ArrayFastSort1([]int{2, 3, 4, 1, 0, 7}); !isArrayEqual(ret, []int{0, 1, 2, 3, 4, 7}) {
		t.Errorf("Test1 error with return: %v", ret)
	}
}

func TestArrayFastSort2(t *testing.T) {
	if ret := ArrayFastSort2([]int{2, 3, 4, 1, 0, 7}); !isArrayEqual(ret, []int{0, 1, 2, 3, 4, 7}) {
		t.Errorf("Test2 error with return: %v", ret)
	}
}

func isArrayEqual(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
