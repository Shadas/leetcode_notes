package _20_Valid_Parentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	if ret := isValid("()[]{}"); !ret {
		t.Error("()[]{} is true, not false.")
	}

	if ret := isValid("()(]{}"); ret {
		t.Error("()(]{} is false, not true.")
	}
}
