package _113_Path_Sum_2

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

func pathSum(root *TreeNode, sum int) [][]int {
	ret := [][]int{}
	cur := []int{}
	pathSumRecursion(root, sum, cur, &ret)
	return ret
}

func pathSumRecursion(root *TreeNode, sum int, cur []int, ret *[][]int) {
	if root == nil {
		return
	}
	cur = append(cur, root.Val)
	if root.Left == nil && root.Right == nil && sum == root.Val {
		cur1 := []int{}
		for _, v := range cur {
			cur1 = append(cur1, v)
		}
		*ret = append(*ret, cur1)
	} else {
		pathSumRecursion(root.Left, sum-root.Val, cur, ret)
		pathSumRecursion(root.Right, sum-root.Val, cur, ret)
	}
	cur = cur[:len(cur)-1]
}
