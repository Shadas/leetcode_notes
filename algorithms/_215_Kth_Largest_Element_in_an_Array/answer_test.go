package _215_Kth_Largest_Element_in_an_Array

import "testing"

func TestFindKthLargest(t *testing.T) {
	var (
		arr    []int
		k, ret int
	)
	arr, k = []int{3, 2, 1, 5, 6, 4}, 2
	ret = findKthLargest(arr, k)
	if ret != 5 {
		t.Errorf("ret is not 5, it's %d", ret)
	}
	arr, k = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4
	ret = findKthLargest(arr, k)
	if ret != 4 {
		t.Errorf("ret is not 4, it's %d", ret)
	}
}
