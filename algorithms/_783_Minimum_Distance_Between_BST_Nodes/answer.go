package _783_Minimum_Distance_Between_BST_Nodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	return minDiffInBSTInOrderNR(root)
}

func minDiffInBSTInOrderNR(root *TreeNode) int {
	var (
		stack []*TreeNode
		node  = root
		pre   *TreeNode
		min   int = math.MaxInt32
	)
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil {
				min = minInt(min, absInt(node.Val, pre.Val))
			}
			pre = node
			node = node.Right
		}
	}
	return min
}

func absInt(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
