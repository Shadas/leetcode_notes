package _2181_Merge_Nodes_in_Between_Zeros

import (
	"testing"

	"github.com/shadas/leetcode_notes/utils/linkedlist"
)

type testCase struct {
	input  *linkedlist.IntListNode
	output *linkedlist.IntListNode
}

func TestMergeNodes(t *testing.T) {
	cases := []testCase{
		{
			input:  linkedlist.GenerateIntLinkedList([]int{0, 3, 1, 0, 4, 5, 2, 0}),
			output: linkedlist.GenerateIntLinkedList([]int{4, 11}),
		},
		{
			input:  linkedlist.GenerateIntLinkedList([]int{0, 1, 0, 3, 0, 2, 2, 0}),
			output: linkedlist.GenerateIntLinkedList([]int{1, 3, 4}),
		},
	}
	for _, c := range cases {
		if x := mergeNodes(c.input); !linkedlist.IsTwoIntLinkedListEqual(x, c.output) {
			t.Errorf("output should be \"%v\" instead of \"%v\" with input=\"%v\"", c.output, x, c.input)
		}
	}
}
