package _671_Second_Minimum_Node_In_a_Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	return findSecondMinimumValueWithRunAll(root)
}

func findSecondMinimumValueWithRunAll(root *TreeNode) int {
	if root == nil {
		return -1
	}
	var (
		min = root.Val
		l   = []*TreeNode{root}
		ret = min
	)
	for len(l) > 0 {
		x := l[len(l)-1]
		l = l[:len(l)-1]
		if ret == min { // 说明还没有
			if x.Val > min {
				ret = x.Val
			}
		} else {
			if x.Val > min && x.Val < ret {
				ret = x.Val
			}
		}
		if x.Left != nil {
			l = append(l, x.Left)
		}
		if x.Right != nil {
			l = append(l, x.Right)
		}
	}
	if ret == min {
		return -1
	}
	return ret
}
