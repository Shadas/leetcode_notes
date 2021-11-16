package _560_Subarray_Sum_Equals_K

import "testing"

func TestSubarraySum(t *testing.T) {
	if count := subarraySum([]int{1, 1, 1}, 2); count != 2 {
		t.Errorf("wrong count with %d", count)
	}
}
