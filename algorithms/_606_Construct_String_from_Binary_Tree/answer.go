package _606_Construct_String_from_Binary_Tree

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	return tree2strPreR(root)
}

func tree2strPreR(root *TreeNode) string {
	if root == nil {
		return ""
	}
	x := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return x
	}
	ret := x + "(" + tree2strPreR(root.Left) + ")"
	if root.Right != nil {
		ret += "(" + tree2strPreR(root.Right) + ")"
	}
	return ret
}
