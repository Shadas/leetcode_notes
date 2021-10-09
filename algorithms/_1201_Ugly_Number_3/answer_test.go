package _1201_Ugly_Number_3

import "testing"

func TestNthUglyNumber(t *testing.T) {
	if ret := nthUglyNumber(1000000000, 2, 217983653, 336916467); ret != 1999999984 {
		t.Errorf("wrong ret with %d", ret)
	}
	if ret := nthUglyNumber(5, 2, 3, 3); ret != 8 {
		t.Errorf("wrong ret with %d", ret)
	}
}
