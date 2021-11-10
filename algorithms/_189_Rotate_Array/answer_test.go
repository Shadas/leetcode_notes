package _189_Rotate_Array

import (
	"testing"

	"github.com/shadas/leetcode_notes/utils/array"
)

func TestRotate(t *testing.T) {
	var (
		nums []int
	)
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	if rotate(nums, 3); !array.IsIntArrayEqual(nums, []int{5, 6, 7, 1, 2, 3, 4}) {
		t.Errorf("wrong ret with %v", nums)
	}
}
