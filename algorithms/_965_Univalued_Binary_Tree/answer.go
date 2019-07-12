package _965_Univalued_Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isUnivalTreeWithValue(root, root.Val)
}

func isUnivalTreeWithValue(root *TreeNode, val int) bool {
	var ret bool
	if root == nil {
		return true
	}
	if root.Val != val {
		return ret
	}
	if root.Left != nil && !isUnivalTreeWithValue(root.Left, val) {
		return ret
	}
	if root.Right != nil && !isUnivalTreeWithValue(root.Right, val) {
		return ret
	}
	return true
}
