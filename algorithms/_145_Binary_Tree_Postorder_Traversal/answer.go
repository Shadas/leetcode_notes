package _145_Binary_Tree_Postorder_Traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	// return postorderTraversalRecursion(root)
	return postorderTraversalUnrecursionWithTwoStack(root)
}

func postorderTraversalUnrecursionWithTwoStack(node *TreeNode) []int {
	var (
		ret  = []int{}
		l    = []*TreeNode{node}
		lret = []*TreeNode{}
	)
	if node == nil {
		return ret
	}
	for len(l) > 0 {
		x := l[len(l)-1]
		l = l[:len(l)-1]
		lret = append(lret, x)
		if x.Left != nil {
			l = append(l, x.Left)
		}
		if x.Right != nil {
			l = append(l, x.Right)
		}
	}
	for len(lret) > 0 {
		x := lret[len(lret)-1]
		lret = lret[:len(lret)-1]
		ret = append(ret, x.Val)
	}
	return ret
}

func postorderTraversalRecursion(node *TreeNode) []int {
	var (
		ret = []int{}
	)
	if node == nil {
		return ret
	}
	if node.Left != nil {
		ret = append(ret, postorderTraversalRecursion(node.Left)...)
	}
	if node.Right != nil {
		ret = append(ret, postorderTraversalRecursion(node.Right)...)
	}
	ret = append(ret, node.Val)
	return ret
}
