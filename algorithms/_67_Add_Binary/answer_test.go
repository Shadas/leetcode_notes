package _67_Add_Binary

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	if ret := addBinary("11", "1"); ret != "100" {
		t.Error("ret not 100, get", ret)
	}
	if ret := addBinary("100", "110010"); ret != "110110" {
		t.Error("ret not 110110, get", ret)
	}
}
