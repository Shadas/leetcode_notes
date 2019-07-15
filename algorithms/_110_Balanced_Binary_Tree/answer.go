package _110_Balanced_Binary_Tree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if d := depth(root); d == -1 {
		return false
	}
	return true
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	var (
		leftDepth, rightDepth int
	)
	if leftDepth = depth(node.Left); leftDepth == -1 {
		return -1
	}
	if rightDepth = depth(node.Right); rightDepth == -1 {
		return -1
	}
	if int(math.Abs(float64(leftDepth-rightDepth))) > 1 {
		return -1
	}
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}
