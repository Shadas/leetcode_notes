package _53_Maximum_Subarray

import (
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	ret := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	if ret != 6 {
		t.Error("not 6, is", ret)
	}
}
