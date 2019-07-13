package _144_Binary_Tree_Preorder_Traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	return preorderTraversalRecursion(root)
}

func preorderTraversalRecursion(n *TreeNode) []int {
	var (
		ret = []int{}
	)
	if n == nil {
		return ret
	}
	ret = append(ret, n.Val)
	if n.Left != nil {
		ret = append(ret, preorderTraversalRecursion(n.Left)...)
	}
	if n.Right != nil {
		ret = append(ret, preorderTraversalRecursion(n.Right)...)
	}
	return ret
}
