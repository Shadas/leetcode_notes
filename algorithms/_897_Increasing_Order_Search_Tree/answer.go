package _897_Increasing_Order_Search_Tree

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	var (
		s    = []*TreeNode{}
		ol   = []int{}
		node = root
		head = &TreeNode{}
	)
	for node != nil || len(s) != 0 {
		for node != nil {
			s = append(s, node)
			node = node.Left
		}
		if len(s) != 0 {
			node = s[len(s)-1]
			s = s[0 : len(s)-1]
			ol = append(ol, node.Val)
			node = node.Right
		}
	}
	fmt.Println(ol)
	// 构建新的二叉树
	tmp := head
	for _, i := range ol {
		tmp.Right = &TreeNode{
			Val: i,
		}
		tmp = tmp.Right
	}
	return head.Right
}
