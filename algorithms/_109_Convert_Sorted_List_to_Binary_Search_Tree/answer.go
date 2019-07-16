package _109_Convert_Sorted_List_to_Binary_Search_Tree

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printList(l *ListNode) {
	vals := []int{}
	tmp := l
	for tmp != nil {
		vals = append(vals, tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println(vals)
}

func sortedListToBST(head *ListNode) *TreeNode {
	return sortedListToBSTRecursion(head)
}

func sortedListToBSTRecursion(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var (
		ret         *TreeNode
		val, l1, l2 = cutList(head)
	)
	ret = &TreeNode{
		Val: val,
	}
	if l1 != nil {
		ret.Left = sortedListToBSTRecursion(l1)
	}
	if l2 != nil {
		ret.Right = sortedListToBSTRecursion(l2)
	}

	return ret
}

func cutList(head *ListNode) (val int, l1, l2 *ListNode) {
	var (
		fast, slow = head, head
	)
	l1 = head
	if head == nil {
		val = 0
		l1 = nil
		l2 = nil
		return
	}
	if head.Next == nil {
		val = head.Val
		l1 = nil
		l2 = nil
		return
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	val = slow.Val
	l2 = slow.Next
	fast = head
	for fast.Next != slow && fast.Next != nil {
		fast = fast.Next
	}
	fast.Next = nil
	return
}
