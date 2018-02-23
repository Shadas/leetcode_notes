package _66_Plus_One

import (
	. "leetcode_notes/utils/array"
	"testing"
)

func TestPlusOne(t *testing.T) {
	if ret := plusOne([]int{8, 9, 9, 9}); !IsIntArrayEqual(ret, []int{9, 0, 0, 0}) {
		t.Errorf("error with return: %v", ret)
	}
}
