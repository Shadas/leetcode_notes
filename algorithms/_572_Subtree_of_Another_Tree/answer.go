package _572_Subtree_of_Another_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	// return isSubtreeUnrecursion(s, t)
	return isSubtreeRecursion(s, t)
}

func isSubtreeUnrecursion(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}
	var (
		l = []*TreeNode{s}
	)
	for len(l) > 0 {
		x := l[len(l)-1]
		l = l[:len(l)-1]
		if x == nil {
			continue
		}
		if isEqual(x, t) {
			return true
		}
		l = append(l, x.Left, x.Right)
	}
	return false
}

func isSubtreeRecursion(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}
	if s == nil {
		return false
	}
	if isEqual(s, t) {
		return true
	}
	return isSubtreeRecursion(s.Left, t) || isSubtreeRecursion(s.Right, t)
}

func isEqual(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s != nil && t == nil {
		return false
	}
	if s == nil && t != nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
}
