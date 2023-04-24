package _409_Longest_Palindrome

import "testing"

type testCase struct {
	input  string
	output int
}

func TestLongestPalindrome(t *testing.T) {
	cases := []testCase{
		{
			input:  "abccccdd",
			output: 7,
		},
		{
			input:  "a",
			output: 1,
		},
		{
			input:  "bb",
			output: 2,
		},
		{
			input:  "ccc",
			output: 3,
		},
	}
	for _, c := range cases {
		if x := longestPalindrome(c.input); x != c.output {
			t.Errorf("output should be \"%v\" instead of \"%v\" with input=\"%v\"", c.output, x, c.input)
		}
	}
}
