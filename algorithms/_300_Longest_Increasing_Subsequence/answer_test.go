package _300_Longest_Increasing_Subsequence

import "testing"

func TestLengthOfLIS(t *testing.T) {
	if ret := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}); ret != 4 {
		t.Errorf("should be 4, wrong length with %d", ret)
	}
}

func TestFindPos(t *testing.T) {
	if pos := findPos([]int{1, 2, 3}, 1); pos != 0 {
		t.Errorf("wrong pos with %d", pos)
	}
	if pos := findPos([]int{1, 2, 3}, 2); pos != 1 {
		t.Errorf("wrong pos with %d", pos)
	}
	if pos := findPos([]int{1, 2, 4}, 3); pos != 2 {
		t.Errorf("wrong pos with %d", pos)
	}
}
