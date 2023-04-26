package _2487_Remove_Nodes_From_Linked_List

import (
	"testing"

	"github.com/shadas/leetcode_notes/utils/linkedlist"
)

type testCase struct {
	input  *linkedlist.IntListNode
	output *linkedlist.IntListNode
}

func TestRemoveNodes(t *testing.T) {
	cases := []testCase{
		{
			input:  linkedlist.GenerateIntLinkedList([]int{5, 2, 13, 3, 8}),
			output: linkedlist.GenerateIntLinkedList([]int{13, 8}),
		},
		{
			input:  linkedlist.GenerateIntLinkedList([]int{1, 1, 1, 1}),
			output: linkedlist.GenerateIntLinkedList([]int{1, 1, 1, 1}),
		},
	}
	for _, c := range cases {
		if x := removeNodes(c.input); !linkedlist.IsTwoIntLinkedListEqual(x, c.output) {
			t.Errorf("output should be \"%v\" instead of \"%v\" with input=\"%v\"", c.output, x, c.input)
		}
	}
}
