package _938_Range_Sum_of_BST

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	var sum int
	if root.Val >= L && root.Left != nil {
		sum += rangeSumBST(root.Left, L, R)
	}
	if root.Val <= R && root.Right != nil {
		sum += rangeSumBST(root.Right, L, R)
	}
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	return sum
}
