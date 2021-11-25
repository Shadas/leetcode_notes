package _647_Palindromic_Substrings

import "testing"

func TestCountSubstrings(t *testing.T) {
	if n := countSubstrings("abc"); n != 3 {
		t.Errorf("wrong count with %d", n)
	}
	if n := countSubstrings("aaa"); n != 6 {
		t.Errorf("wrong count with %d", n)
	}
}
