package _530_Minimum_Absolute_Difference_in_BST

import "testing"

func TestGetMinimumDifference(t *testing.T) {
	var (
		root *TreeNode
		ret  int
	)
	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 48, Left: &TreeNode{Val: 12}, Right: &TreeNode{Val: 49}}}
	if ret = getMinimumDifference(root); ret != 1 {
		t.Errorf("wrong ret with %d", ret)
	}
	root = &TreeNode{Val: 236, Left: &TreeNode{Val: 104, Right: &TreeNode{Val: 227}}, Right: &TreeNode{Val: 701, Right: &TreeNode{Val: 911}}}
	if ret = getMinimumDifference(root); ret != 9 {
		t.Errorf("wrong ret with %d", ret)
	}
	root = &TreeNode{Val: 2, Right: &TreeNode{Val: 4443, Left: &TreeNode{Val: 1329, Right: &TreeNode{Val: 2917, Right: &TreeNode{Val: 4406}}}}}
	if ret = getMinimumDifference(root); ret != 37 {
		t.Errorf("wrong ret with %d", ret)
	}
}
