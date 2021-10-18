package _144_Binary_Tree_Preorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// return preorderTraversalRecursion(root)
	return preorderTraversalUnrecursion(root)
}

func preorderTraversalUnrecursion(root *TreeNode) []int {
	var (
		ret = []int{}
		l   = []*TreeNode{root}
	)
	if root == nil {
		return ret
	}
	for len(l) > 0 {
		node := l[len(l)-1]
		ret = append(ret, node.Val)
		l = l[:len(l)-1]
		if node.Right != nil {
			l = append(l, node.Right)
		}
		if node.Left != nil {
			l = append(l, node.Left)
		}
	}
	return ret
}

func preorderTraversalRecursion(n *TreeNode) []int {
	var (
		ret = []int{}
	)
	if n == nil {
		return ret
	}
	ret = append(ret, n.Val)
	if n.Left != nil {
		ret = append(ret, preorderTraversalRecursion(n.Left)...)
	}
	if n.Right != nil {
		ret = append(ret, preorderTraversalRecursion(n.Right)...)
	}
	return ret
}
