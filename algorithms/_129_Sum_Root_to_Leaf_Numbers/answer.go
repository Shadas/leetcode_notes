package _129_Sum_Root_to_Leaf_Numbers

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var (
		traces [][]int
		sum    int
	)
	findTrace(root, []int{}, &traces)
	for _, trace := range traces {
		sum += transTrace2Num(trace)
	}
	return sum
}

func transTrace2Num(t []int) (num int) {
	for i := 0; i < len(t); i++ {
		num += t[i] * int(math.Pow(10, float64(len(t)-i-1)))
	}
	return
}

func findTrace(node *TreeNode, trace []int, ret *[][]int) {
	trace = append(trace, node.Val)
	if node.Left == nil && node.Right == nil {
		var tmpTrace []int
		for _, v := range trace {
			tmpTrace = append(tmpTrace, v)
		}
		*ret = append(*ret, tmpTrace)
		return
	}
	if node.Left != nil {
		findTrace(node.Left, trace, ret)
	}
	if node.Right != nil {
		findTrace(node.Right, trace, ret)
	}
}
