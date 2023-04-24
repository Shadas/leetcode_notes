package _334_Increasing_Triplet_Subsequence

import (
	"testing"
)

type testCase struct {
	input  []int
	output bool
}

func TestIncreasingTriplet(t *testing.T) {
	cases := []testCase{
		{
			input:  []int{1, 2, 3, 4, 5},
			output: true,
		},
		{
			input:  []int{5, 4, 3, 2, 1},
			output: false,
		},
		{
			input:  []int{2, 1, 5, 0, 4, 6},
			output: true,
		},
	}
	for _, c := range cases {
		if x := increasingTriplet(c.input); x != c.output {
			t.Errorf("output should be \"%v\" instead of \"%v\" with input=\"%v\"", c.output, x, c.input)
		}
	}
}
