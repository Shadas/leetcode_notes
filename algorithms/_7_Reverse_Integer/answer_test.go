package _7_Reverse_Integer

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if ret := reverse(-12); ret != -21 {
		t.Error("not -21 with -12.")
	}
	if ret := reverse(1200); ret != 21 {
		t.Error("not 21 with 1200.")
	}
	if ret := reverse(1534236469); ret != 0 {
		t.Error("not 0 with 1534236469.")
	}
	if ret := reverse(-2147483412); ret != -2143847412 {
		t.Error("not -2143847412 with -2147483412.")
	}
	if ret := reverse(1463847412); ret != 2147483641 {
		t.Error("not 2147483641 with 1463847412.")
	}
}
