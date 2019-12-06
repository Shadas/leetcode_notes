package _701_Insert_into_a_Binary_Search_Tree

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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	var (
		node = root
	)
	for {
		if node.Val < val {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
				return root
			} else {
				node = node.Right
			}
		}
		if node.Val > val {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				return root
			} else {
				node = node.Left
			}
		}
	}
	return root
}
