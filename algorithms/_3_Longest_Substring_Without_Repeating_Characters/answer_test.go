package _3_Longest_Substring_Without_Repeating_Characters

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	input1 := "abcabcab"

	ret1 := lengthOfLongestSubstring(input1)
	if ret1 != 3 {
		t.Errorf("\"%v\" 's ret is %v not %v", input1, ret1, 3)
	}

	input2 := "bbbbbbbb"

	ret2 := lengthOfLongestSubstring(input2)
	if ret2 != 1 {
		t.Errorf("\"%v\" 's ret is %v not %v", input2, ret2, 3)
	}

	input3 := "pwwkew"

	ret3 := lengthOfLongestSubstring(input3)
	if ret3 != 3 {
		t.Errorf("\"%v\" 's ret is %v not %v", input3, ret3, 3)
	}

	input4 := "au"

	ret4 := lengthOfLongestSubstring(input4)
	if ret4 != 2 {
		t.Errorf("\"%v\" 's ret is %v not %v", input4, ret4, 2)
	}

	input5 := "c"

	ret5 := lengthOfLongestSubstring(input5)
	if ret5 != 1 {
		t.Errorf("\"%v\" 's ret is %v not %v", input5, ret5, 1)
	}

	input6 := "abcdef"

	ret6 := lengthOfLongestSubstring(input6)
	if ret6 != 6 {
		t.Errorf("\"%v\" 's ret is %v not %v", input6, ret6, 6)
	}
}
