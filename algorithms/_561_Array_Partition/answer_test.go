package _561_Array_Partition

import "testing"

type testCase struct {
	input  []int
	output int
}

func TestArrayPairSum(t *testing.T) {
	cases := []testCase{
		{
			input:  []int{1, 4, 3, 2},
			output: 4,
		},
		{
			input:  []int{6, 2, 6, 5, 1, 2},
			output: 9,
		},
	}
	for _, c := range cases {
		if x := arrayPairSum(c.input); x != c.output {
			t.Errorf("output should be \"%v\" instead of \"%v\" with input=\"%v\"", c.output, x, c.input)
		}
	}
}
