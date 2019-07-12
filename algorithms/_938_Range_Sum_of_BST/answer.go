package _938_Range_Sum_of_BST

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	// return rangeSumBSTUnrecursion(root, L, R)
	return rangeSumBSTRecursion(root, L, R)
}

func rangeSumBSTUnrecursion(root *TreeNode, L int, R int) int {
	var (
		l   []*TreeNode = []*TreeNode{}
		sum int
	)
	l = append(l, root)
	for len(l) > 0 {
		x := l[len(l)-1]
		l = l[:len(l)-1]
		if x == nil {
			continue
		}
		if x.Val < L {
			if x.Right != nil {
				l = append(l, x.Right)
			}
		} else if x.Val > R {
			if x.Left != nil {
				l = append(l, x.Left)
			}
		} else {
			sum += x.Val
			l = append(l, x.Left, x.Right)
		}
	}
	return sum
}

func rangeSumBSTRecursion(root *TreeNode, L int, R int) int {
	var sum int
	if root.Val >= L && root.Left != nil {
		sum += rangeSumBST(root.Left, L, R)
	}
	if root.Val <= R && root.Right != nil {
		sum += rangeSumBST(root.Right, L, R)
	}
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	return sum
}
