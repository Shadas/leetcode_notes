package _111_Minimum_Depth_of_Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	return minDepthRecursively(root)
}

// recursively
func minDepthRecursively(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	} else if root.Right == nil {
		return 1 + minDepth(root.Left)
	} else {
		return 1 + min(minDepthRecursively(root.Left), minDepthRecursively(root.Right))
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
