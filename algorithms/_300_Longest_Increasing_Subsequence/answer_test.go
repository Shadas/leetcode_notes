package _300_Longest_Increasing_Subsequence

import "testing"

func TestLengthOfLIS(t *testing.T) {
	if ret := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}); ret != 4 {
		t.Errorf("should be 4, wrong length with %d", ret)
	}
}
