package _862_Shortest_Subarray_with_Sum_at_Least_K

import "testing"

func TestShortestSubarray(t *testing.T) {
	var res int
	if res = shortestSubarray([]int{1}, 1); res != 1 {
		t.Errorf("wrong res=%d", res)
	}
}
