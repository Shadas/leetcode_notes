package _28_Implement_StrStr

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	if ret := strStr("hello", "ll"); ret != 2 {
		t.Error("Error not 2, get", ret)
	}
	if ret := strStr("aaaaa", "bba"); ret != -1 {
		t.Error("Error not -1, get", ret)
	}
	if ret := strStr("mississippi", "mississippi"); ret != 0 {
		t.Error("Error not 0, get", ret)
	}
}
