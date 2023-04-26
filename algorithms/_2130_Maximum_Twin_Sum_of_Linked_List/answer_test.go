package _2130_Maximum_Twin_Sum_of_Linked_List

import (
	"testing"

	"github.com/shadas/leetcode_notes/utils/linkedlist"
)

type testCase struct {
	input  *linkedlist.IntListNode
	output int
}

func TestPairSum(t *testing.T) {
	cases := []testCase{
		{
			input:  linkedlist.GenerateIntLinkedList([]int{5, 4, 2, 1}),
			output: 6,
		},
		{
			input:  linkedlist.GenerateIntLinkedList([]int{4, 2, 2, 3}),
			output: 7,
		},
		{
			input:  linkedlist.GenerateIntLinkedList([]int{1, 100000}),
			output: 100001,
		},
	}
	for _, c := range cases {
		if x := pairSum(c.input); x != c.output {
			t.Errorf("output should be \"%v\" instead of \"%v\" with input=\"%v\"", c.output, x, c.input)
		}
	}
}
