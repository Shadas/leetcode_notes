package _226_Invert_Binary_Tree

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

func invertTree(root *TreeNode) *TreeNode {
	// return invertTreeRecursive(root)
	return invertTreeIterative(root)
}

// 递归
func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	tmpLeft := root.Left
	root.Left = invertTreeRecursive(root.Right)
	root.Right = invertTreeRecursive(tmpLeft)
	return root
}

// 迭代
func invertTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var (
		level = []*TreeNode{root}
	)
	for len(level) != 0 {
		currLevel := make([]*TreeNode, len(level))
		copy(currLevel, level)
		level = level[0:0]
		for _, n := range currLevel {
			n.Left, n.Right = n.Right, n.Left
			if n.Left != nil {
				level = append(level, n.Left)
			}
			if n.Right != nil {
				level = append(level, n.Right)
			}
		}
	}
	return root
}
