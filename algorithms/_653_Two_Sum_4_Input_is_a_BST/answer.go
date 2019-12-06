package _653_Two_Sum_4_Input_is_a_BST

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

func findTarget(root *TreeNode, k int) bool {
	var (
		m = make(map[int]*TreeNode)
	)
	runTree(root, m)
	return findTree(root, k, m)
}

func findTree(root *TreeNode, target int, m map[int]*TreeNode) bool {
	if root == nil {
		return false
	}
	if node, ok := m[target-root.Val]; ok && node != root {
		return true
	}
	if left := findTree(root.Left, target, m); left {
		return true
	}
	if right := findTree(root.Right, target, m); right {
		return true
	}
	return false
}

func runTree(root *TreeNode, m map[int]*TreeNode) {
	if root == nil {
		return
	}
	m[root.Val] = root
	runTree(root.Left, m)
	runTree(root.Right, m)
	return
}
