package _437_Path_Sum_3

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

func pathSum(root *TreeNode, sum int) int {
	preSum := make(map[int]int)
	preSum[0] = 1
	return countPath(root, 0, sum, preSum)
}

func countPath(root *TreeNode, currSum int, target int, preSum map[int]int) int {
	if root == nil {
		return 0
	}
	currSum += root.Val
	var res int
	if v, ok := preSum[currSum-target]; ok {
		res = v
	} else {
		res = 0
	}
	var newv int
	if v, ok := preSum[currSum]; ok {
		newv = v + 1
	} else {
		newv = 0 + 1
	}
	preSum[currSum] = newv
	res += countPath(root.Left, currSum, target, preSum) + countPath(root.Right, currSum, target, preSum)
	preSum[currSum] = preSum[currSum] - 1
	return res
}
