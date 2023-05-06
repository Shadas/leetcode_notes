package _606_Construct_String_from_Binary_Tree

import (
	"testing"
)

type testCase struct {
	input  *TreeNode
	output string
}

func TestTree2str(t *testing.T) {
	cases := []testCase{
		{
			input:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
			output: "1(2(4))(3)",
		},
	}
	for _, c := range cases {
		if x := tree2str(c.input); x != c.output {
			t.Errorf("output should be \"%v\" instead of \"%v\" with input=\"%v\"", c.output, x, c.input)
		}
	}
}
