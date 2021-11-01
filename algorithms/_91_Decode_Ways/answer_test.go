package _91_Decode_Ways

import "testing"

func TestNumDecodings(t *testing.T) {
	if ret := numDecodings("12"); ret != 2 {
		t.Errorf("wrong answer with %d", ret)
	}
	if ret := numDecodings("226"); ret != 3 {
		t.Errorf("wrong answer with %d", ret)
	}
	if ret := numDecodings("0"); ret != 0 {
		t.Errorf("wrong answer with %d", ret)
	}
}
