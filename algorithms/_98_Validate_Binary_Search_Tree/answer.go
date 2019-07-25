package _98_Validate_Binary_Search_Tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {

	return isValidBSTRecursion(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTRecursion(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}

	return isValidBSTRecursion(node.Left, min, node.Val) && isValidBSTRecursion(node.Right, node.Val, max)
}
