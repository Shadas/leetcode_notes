package _538_Convert_BST_to_Greater_Tree

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

func convertBST(root *TreeNode) *TreeNode {
	var sum int = 0
	return postOrder(root, &sum)
}

func postOrder(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return nil
	}
	postOrder(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	postOrder(root.Left, sum)
	return root
}
