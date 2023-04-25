package _581_Shortest_Unsorted_Continuous_Subarray

import "testing"

type testCase struct {
	input  []int
	output int
}

func TestFindUnsortedSubarray(t *testing.T) {
	cases := []testCase{
		{
			input:  []int{2, 6, 4, 8, 10, 9, 15},
			output: 5,
		},
		{
			input:  []int{1, 2, 3, 4},
			output: 0,
		},
		{
			input:  []int{4, 3, 2, 1},
			output: 4,
		},
		{
			input:  []int{1},
			output: 0,
		},
	}
	for _, c := range cases {
		if x := findUnsortedSubarray(c.input); x != c.output {
			t.Errorf("output should be \"%v\" instead of \"%v\" with input=\"%v\"", c.output, x, c.input)
		}
	}
}
