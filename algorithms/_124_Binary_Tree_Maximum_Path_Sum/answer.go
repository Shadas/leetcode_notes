package _124_Binary_Tree_Maximum_Path_Sum

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	var ret = math.MinInt32
	calc(root, &ret)
	return ret
}

func calc(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	lret := int(math.Max(float64(calc(node.Left, max)), 0))
	rret := int(math.Max(float64(calc(node.Right, max)), 0))
	*max = int(math.Max(float64(*max), float64(lret+rret+node.Val)))
	return int(math.Max(float64(lret+node.Val), float64(rret+node.Val)))
}
