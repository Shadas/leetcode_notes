package _988_Smallest_String_Starting_From_Leaf

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func smallestFromLeaf(root *TreeNode) string {
	var (
		traces   = [][]int{}
		minTrace []int
	)
	if root == nil {
		return ""
	}
	findTraces(root, []int{}, &traces)
	// 找最小的trace
	minTrace = traces[0]
	for _, trace := range traces {
		minTrace = minTraces(minTrace, trace)
	}
	return transNumbersToStr(minTrace)
}

func minTraces(t1, t2 []int) []int {
	for i := 0; ; i++ {
		if len(t1) < i+1 {
			return t1
		}
		if len(t2) < i+1 {
			return t2
		}
		if t1[i] > t2[i] {
			return t2
		} else if t1[i] < t2[i] {
			return t1
		}
	}
}

func transNumbersToStr(nums []int) (out string) {
	for _, v := range nums {
		out += fmt.Sprintf("%c", v+97)
	}
	return
}

func findTraces(node *TreeNode, trace []int, ret *[][]int) {
	trace = append(trace, node.Val)
	if node.Left == nil && node.Right == nil {
		nTrace := make([]int, 0, len(trace))
		for i := len(trace) - 1; i >= 0; i-- {
			nTrace = append(nTrace, trace[i])
		}
		*ret = append(*ret, nTrace)
		return
	}
	if node.Left != nil {
		findTraces(node.Left, trace, ret)
	}
	if node.Right != nil {
		findTraces(node.Right, trace, ret)
	}
}
