package _2095_Delete_the_Middle_Node_of_a_Linked_List

import (
	"testing"

	"github.com/shadas/leetcode_notes/utils/linkedlist"
)

type testCase struct {
	input  *linkedlist.IntListNode
	output *linkedlist.IntListNode
}

func TestDeleteMiddle(t *testing.T) {
	cases := []testCase{
		{
			input:  linkedlist.GenerateIntLinkedList([]int{1, 3, 4, 7, 1, 2, 6}),
			output: linkedlist.GenerateIntLinkedList([]int{1, 3, 4, 1, 2, 6}),
		},
	}
	for _, c := range cases {
		if x := deleteMiddle(c.input); !linkedlist.IsTwoIntLinkedListEqual(x, c.output) {
			t.Errorf("output should be \"%v\" instead of \"%v\" with input=\"%v\"", c.output, x, c.input)
		}
	}
}
