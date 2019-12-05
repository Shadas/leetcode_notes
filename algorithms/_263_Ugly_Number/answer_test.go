package _263_Ugly_Number

import "testing"

func TestIsUgly(t *testing.T) {
	if ret := isUgly(1); ret != true {
		t.Errorf("%d ret should be %v", 1, true)
	}
	if ret := isUgly(6); ret != true {
		t.Errorf("%d ret should be %v", 6, true)
	}
}
