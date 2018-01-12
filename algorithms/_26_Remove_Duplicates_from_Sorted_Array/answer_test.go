package _26_Remove_Duplicates_from_Sorted_Array

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	if ret := removeDuplicates([]int{1, 1, 2}); ret != 2 {
		t.Error("ret not 2 with:", ret)
	}
}
