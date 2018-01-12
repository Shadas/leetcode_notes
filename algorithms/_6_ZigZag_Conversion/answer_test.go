package _6_ZigZag_Conversion

import (
	"testing"
)

func TestConvert(t *testing.T) {
	if ret := convert("PAYPALISHIRING", 3); ret != "PAHNAPLSIIGYIR" {
		t.Error("err ret with test1.")
	}
}
