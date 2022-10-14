package _543_Diameter_of_Binary_Tree

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	r := &TreeNode{
		Left: &TreeNode{
			Value: 1,
		},
		Value: 2,
	}
	ret := diameterOfBinaryTree(r)
	t.Log(ret)
}
