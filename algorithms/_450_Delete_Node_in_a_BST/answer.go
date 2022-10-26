package _450_Delete_Node_in_a_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	pre := &TreeNode{Left: root}
	deleteNodeR(root, pre, true, key)
	return pre.Left
}

func deleteNodeR(root, pre *TreeNode, isLeft bool, key int) {
	if root == nil {
		return
	}
	if root.Val > key {
		deleteNodeR(root.Left, root, true, key)
		return
	}
	if root.Val < key {
		deleteNodeR(root.Right, root, false, key)
		return
	}
	if root.Val == key { // delete
		var tmp *TreeNode
		if root.Left == nil {
			tmp = root.Right
		} else if root.Right == nil {
			tmp = root.Left
		} else {
			tmp = buildBST(root.Left, root.Right)
		}
		if isLeft {
			pre.Left = tmp
		} else {
			pre.Right = tmp
		}
	}
}

func buildBST(left, right *TreeNode) *TreeNode {
	n := left
	lr := n.Right
	n.Right = right
	for right.Left != nil {
		right = right.Left
	}
	right.Left = lr
	return n
}
