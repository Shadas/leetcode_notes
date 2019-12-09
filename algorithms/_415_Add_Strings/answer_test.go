package _415_Add_Strings

import "testing"

func TestAddStrings(t *testing.T) {
	if ret := addStrings("10", "2"); ret != "12" {
		t.Errorf("wrong with %s", ret)
	}
	if ret := addStrings("99", "1"); ret != "100" {
		t.Errorf("wrong with %s", ret)
	}
	if ret := addStrings("199", "199"); ret != "398" {
		t.Errorf("wrong with %s", ret)
	}
}
