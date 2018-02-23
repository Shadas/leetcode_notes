package _38_Count_and_Say

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	if ret := countAndSay(1); ret != "1" {
		t.Error("ret not 1, get", ret)
	}
	if ret := countAndSay(2); ret != "11" {
		t.Error("ret not 11, get", ret)
	}
	if ret := countAndSay(3); ret != "21" {
		t.Error("ret not 21, get", ret)
	}
}
