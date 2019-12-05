package _199_Binary_Tree_Right_Side_View

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
	Right *TreeNod
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		level = []*TreeNode{root}
		ret   = []int{}
	)
	for len(level) != 0 {
		ret = append(ret, level[len(level)-1].Val)
		nextLevel := []*TreeNode{}
		for _, n := range level {
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
			level = nextLevel
		}
	}
	return ret
}
