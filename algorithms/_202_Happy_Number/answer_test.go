package _202_Happy_Number

import "testing"

func TestIsHappy(t *testing.T) {
	if isHappy(19) != true {
		t.Log("19 failed")
	}
	if isHappy(20) != true {
		t.Log("20 failed")
	}
}