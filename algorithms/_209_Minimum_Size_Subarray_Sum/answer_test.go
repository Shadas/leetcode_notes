package _209_Minimum_Size_Subarray_Sum

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	//if ret := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}); ret != 2 {
	//	t.Errorf("wrong ret with %d", ret)
	//}
	if ret := minSubArrayLen(11, []int{1, 2, 3, 4, 5}); ret != 3 {
		t.Errorf("wrong ret with %d", ret)
	}
}
