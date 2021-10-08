package _515_Find_Largest_Value_in_Each_Tree_Row

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		level, tmp []*TreeNode

		max int
		ret []int
	)
	level = append(level, root)
	for len(level) != 0 {
		max = math.MinInt32
		for _, node := range level {
			val := node.Val
			if val > max {
				max = val
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		level = tmp
		tmp = []*TreeNode{}
		ret = append(ret, max)
	}
	return ret
}
