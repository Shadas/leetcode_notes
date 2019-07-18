package _94_Binary_Tree_Inorder_Traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	// return inorderTranversalRecursion(root)
	// return inorderTranversalUnrecursion(root)
	return inorderTranversalUnrecursionWithStack(root)
}

func inorderTranversalUnrecursionWithStack(n *TreeNode) []int {
	var (
		s    = []*TreeNode{}
		ret  = []int{}
		node = n
	)
	for node != nil || len(s) > 0 {
		for node != nil {
			s = append(s, node)
			node = node.Left
		}
		if len(s) > 0 {
			node = s[len(s)-1]
			s = s[:len(s)-1]
			ret = append(ret, node.Val)
			node = node.Right
		}
	}
	return ret
}

func inorderTranversalUnrecursion(n *TreeNode) []int {
	var (
		l    = []*TreeNode{}
		node = n
		ret  = []int{}
	)
	for node != nil || len(l) > 0 {
		if node != nil {
			l = append(l, node)
			node = node.Left
		} else {
			node = l[len(l)-1]
			l = l[:len(l)-1]
			ret = append(ret, node.Val)
			node = node.Right
		}
	}
	return ret
}

func inorderTranversalRecursion(n *TreeNode) []int {
	if n == nil {
		return []int{}
	}
	var ret []int
	if n.Left != nil {
		ret = inorderTranversalRecursion(n.Left)
	}
	ret = append(ret, n.Val)
	if n.Right != nil {
		ret = append(ret, inorderTranversalRecursion(n.Right)...)
	}
	return ret
}
