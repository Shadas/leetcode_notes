package _239_Sliding_Window_Maximum

import (
	"testing"

	"github.com/shadas/leetcode_notes/utils/array"
)

func TestMaxSlidingWindow(t *testing.T) {
	if ret := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3); !array.IsIntArrayEqual([]int{3, 3, 5, 5, 6, 7}, ret) {
		t.Errorf("wrong ret with %v", ret)
	}
	if ret := maxSlidingWindow([]int{1, -1}, 1); !array.IsIntArrayEqual([]int{1, -1}, ret) {
		t.Errorf("wrong ret with %v", ret)
	}
	if ret := maxSlidingWindow([]int{7, 2, 4}, 2); !array.IsIntArrayEqual([]int{7, 4}, ret) {
		t.Errorf("wrong ret with %v", ret)
	}
}
