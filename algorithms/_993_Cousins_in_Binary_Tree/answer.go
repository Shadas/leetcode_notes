package _993_Cousins_in_Binary_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	return isCousinsBFS(root, x, y)
}

func isCousinsBFS(root *TreeNode, x int, y int) bool {
	var (
		xp, yp *TreeNode
		xd, yd int
		s      = []*TreeNode{root}
		nowDep = -1
	)
	if root == nil {
		return false
	}
	for len(s) > 0 {
		nowDep++
		currFloor := s
		cache := []*TreeNode{}
		for _, n := range currFloor {
			if (n.Left != nil && n.Left.Val == x) || (n.Right != nil && n.Right.Val == x) {
				xp = n
				xd = nowDep + 1
			}
			if (n.Left != nil && n.Left.Val == y) || (n.Right != nil && n.Right.Val == y) {
				yp = n
				yd = nowDep + 1
			}
			if xp != nil && yp != nil {
				if xp != yp && xd == yd {
					return true
				} else {
					return false
				}
			}
			if n.Left != nil {
				cache = append(cache, n.Left)
			}
			if n.Right != nil {
				cache = append(cache, n.Right)
			}
		}
		s = cache
	}
	return false
}
