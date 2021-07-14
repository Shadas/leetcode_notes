package _1367_Linked_List_in_Binary_Tree

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	return checkCurrNode(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func checkCurrNode(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val != root.Val {
		return false
	}
	return checkCurrNode(head.Next, root.Left) || checkCurrNode(head.Next, root.Right)
}
