package _872_Leaf_Similar_Trees

import "fmt"

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

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var l1, l2 []int
	getLeaves(root1, &l1)
	getLeaves(root2, &l2)
	fmt.Println(l1)
	fmt.Println(l2)
	if len(l1) != len(l2) {
		return false
	}
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func getLeaves(node *TreeNode, l *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*l = append(*l, node.Val)
		return
	}
	getLeaves(node.Left, l)
	getLeaves(node.Right, l)
	return
}
