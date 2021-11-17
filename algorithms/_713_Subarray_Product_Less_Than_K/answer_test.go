package _713_Subarray_Product_Less_Than_K

import "testing"

func TestNumSubarrayProductLessThanK(t *testing.T) {
	if count := numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100); count != 8 {
		t.Errorf("wrong count with %d", count)
	}
	if count := numSubarrayProductLessThanK([]int{1, 2, 3}, 0); count != 0 {
		t.Errorf("wrong count with %d", count)
	}
}
