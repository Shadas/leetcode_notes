package _264_Ugly_Number_2

import "testing"

func TestNthUglyNumber(t *testing.T) {
	if ret := nthUglyNumber(10); ret != 12 {
		t.Errorf("ret of %d should be %d, rather than %d", 10, 12, ret)
	}
}
