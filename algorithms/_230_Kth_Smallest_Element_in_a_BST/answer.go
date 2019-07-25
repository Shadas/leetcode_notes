package _230_Kth_Smallest_Element_in_a_BST

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var (
		s     = []*TreeNode{}
		node  = root
		count int
	)
	for node != nil || len(s) > 0 {
		for node != nil {
			s = append(s, node)
			node = node.Left
		}
		if len(s) > 0 {
			node = s[len(s)-1]
			s = s[:len(s)-1]
			count++
			if count == k {
				return node.Val
			}
			node = node.Right
		}
	}
	return -1
}
