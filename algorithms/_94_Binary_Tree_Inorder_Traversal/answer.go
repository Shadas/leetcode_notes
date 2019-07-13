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
	return inorderTranversalUnrecursion(root)
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
