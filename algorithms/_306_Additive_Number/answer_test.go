package _306_Additive_Number

import "testing"

func TestIsAdditiveNumber(t *testing.T) {
	if !isAdditiveNumber("112358") {
		t.Errorf("should be true")
	}
	if !isAdditiveNumber("199100199") {
		t.Errorf("should be true")
	}
	if isAdditiveNumber("1023") {
		t.Errorf("should be false")
	}
	if isAdditiveNumber("1203") {
		t.Errorf("should be false")
	}
	if !isAdditiveNumber("121474836472147483648") {
		t.Errorf("should be true")
	}
}
