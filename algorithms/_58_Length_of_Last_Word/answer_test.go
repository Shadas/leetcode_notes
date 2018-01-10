package _58_Length_of_Last_Word

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	if ret := lengthOfLastWord("Hello world"); ret != 5 {
		t.Error("not 5")
	}
}
