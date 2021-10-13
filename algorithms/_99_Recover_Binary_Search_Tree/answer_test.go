package _99_Recover_Binary_Search_Tree

import (
	"fmt"
	"testing"
)

func TestRecoverTree(t *testing.T) {
	var (
		root *TreeNode
	)
	root = &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}}
	recoverTree(root)
	fmt.Println(root.Val, root.Left.Val, root.Left.Right.Val)

	root = &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}}
	recoverTree(root)
	fmt.Println(root.Val, root.Left.Val, root.Right.Val, root.Right.Left.Val)
}
