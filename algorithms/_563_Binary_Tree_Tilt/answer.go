package _563_Binary_Tree_Tilt

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

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sum int
	funcRecursive(root, &sum)
	return sum
}

func funcRecursive(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	left := funcRecursive(root.Left, sum)
	right := funcRecursive(root.Right, sum)
	tmp := 0
	if left >= right {
		tmp = left - right
	} else {
		tmp = right - left
	}
	*sum += tmp
	return left + right + root.Val
}
