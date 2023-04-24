package _316_Remove_Duplicate_Letters

import (
	"strings"
	"testing"
)

type testCase struct {
	input  string
	output string
}

func TestRemoveDuplicateLetters(t *testing.T) {
	cases := []testCase{
		{
			input:  "bcabc",
			output: "abc",
		},
		{
			input:  "cbacdcbc",
			output: "acdb",
		},
	}
	for _, c := range cases {
		if x := removeDuplicateLetters(c.input); !strings.EqualFold(x, c.output) {
			t.Errorf("output should be \"%s\" instead of \"%s\" with input=\"%s\"", c.output, x, c.input)
		}
	}
}
