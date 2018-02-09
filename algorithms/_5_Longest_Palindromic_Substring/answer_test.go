package _5_Longest_Palindromic_Substring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	if ret := longestPalindrome("babad"); ret != "bab" && ret != "aba" {
		t.Error("not \"bab\" or \"aba\" with \"babad\", get:", ret)
	}
	if ret := longestPalindrome("cbbd"); ret != "bb" {
		t.Error("not \"bb\" with \"cbbd\", get:", ret)
	}
}
