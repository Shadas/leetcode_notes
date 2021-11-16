package _523_Continuous_Subarray_Sum

import "testing"

func TestCheckSubarraySum(t *testing.T) {
	if !checkSubarraySum([]int{23, 2, 6, 4, 7}, 6) {
		t.Error("should be true")
	}

	if !checkSubarraySum([]int{23, 2, 4, 6, 6}, 7) {
		t.Error("should be true")
	}

	if checkSubarraySum([]int{1, 0}, 2) {
		t.Error("should be false")
	}

	if checkSubarraySum([]int{1, 2, 12}, 6) {
		t.Error("should be false")
	}

	if !checkSubarraySum([]int{5, 0, 0, 0}, 3) {
		t.Error("should be true")
	}

	if !checkSubarraySum([]int{1, 1}, 1) {
		t.Error("should be true")
	}

	if !checkSubarraySum([]int{0, 0}, 1) {
		t.Error("should be true")
	}
}
