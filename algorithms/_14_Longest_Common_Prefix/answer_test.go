package _14_Longest_Common_Prefix

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	if ret := longestCommonPrefix([]string{"a", "ab"}); ret != "a" {
		t.Error("not a with test1.")
	}
}
