package _530_Minimum_Absolute_Difference_in_BST

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	return getMinimumDifferenceInOrderR(root)
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

// 解法一：中序遍历，递归实现
func getMinimumDifferenceInOrderR(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := math.MaxInt32
	dfs(root, nil, &min)
	return min
}

func dfs(root *TreeNode, pre *TreeNode, min *int) (newPre *TreeNode) {
	if root.Left != nil {
		pre = dfs(root.Left, pre, min)
	}

	if pre != nil {
		//fmt.Println(pre.Val, root.Val, absInt(root.Val, pre.Val))
		*min = minInt(*min, absInt(root.Val, pre.Val))
	}

	pre, newPre = root, root
	if root.Right != nil {
		newPre = dfs(root.Right, pre, min)
	}
	return
}

// 解法二：中序遍历，非递归实现
func getMinimumDifferenceInOrderNR(root *TreeNode) int {
	return 0
}
