package _112_Path_Sum

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

func hasPathSum(root *TreeNode, sum int) bool {
	cur := []int{}
	return pathSumRecursion(root, sum, cur)
}

func pathSumRecursion(root *TreeNode, sum int, cur []int) bool {
	if root == nil {
		return false
	}
	cur = append(cur, root.Val)
	if root.Left == nil && root.Right == nil && sum == root.Val {
		return true
	}
	if pathSumRecursion(root.Left, sum-root.Val, cur) || pathSumRecursion(root.Right, sum-root.Val, cur) {
		return true
	}
	cur = cur[:len(cur)-1]
	return false
}
