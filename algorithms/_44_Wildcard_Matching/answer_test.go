package _44_Wildcard_Matching

import "testing"

func TestIsMatch(t *testing.T) {
	if isMatch("aa", "a") {
		t.Error("aa and a should be false")
	}
	if !isMatch("aa", "*") {
		t.Error("aa and * should be true")
	}
	if isMatch("cb", "?a") {
		t.Error("cb and ?a should be false")
	}
}
