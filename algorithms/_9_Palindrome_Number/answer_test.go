package _9_Palindrome_Number

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if ok := isPalindrome(0); !ok {
		t.Error("false with 0.")
	}
	if ok := isPalindrome(10); ok {
		t.Error("true with 10.")
	}
	if ok := isPalindrome(12321); !ok {
		t.Error("false with 12321.")
	}
}
