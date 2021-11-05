package _76_Minimum_Window_Substring

import "testing"

func TestIsValid(t *testing.T) {
	if b := isValid("BVCA", "BC"); !b {
		t.Error("should be true")
	}
	if b := isValid("BVCA", "BCV"); !b {
		t.Error("should be true")
	}
}

func TestMinWindow(t *testing.T) {
	if str := minWindow("ADOBECODEBANC", "ABC"); str != "BANC" {
		t.Errorf("wrong str with %s", str)
	}
	if str := minWindow("a", "a"); str != "a" {
		t.Errorf("wrong str with %s", str)
	}
	if str := minWindow("a", "aa"); str != "" {
		t.Errorf("wrong str with %s", str)
	}
	if str := minWindow("a", "b"); str != "" {
		t.Errorf("wrong str with %s", str)
	}

}
