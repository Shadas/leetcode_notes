package _2487_Remove_Nodes_From_Linked_List

import "github.com/shadas/leetcode_notes/utils/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *linkedlist.IntListNode) *linkedlist.IntListNode {
	return removeNodesStack(head)
}

func removeNodesStack(head *linkedlist.IntListNode) *linkedlist.IntListNode {
	var stack []int
	for head != nil {
		for len(stack) != 0 && head.Val > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, head.Val)
		head = head.Next
	}
	x := &linkedlist.IntListNode{}
	y := x
	for _, i := range stack {
		x.Next = &linkedlist.IntListNode{
			Val:  i,
			Next: nil,
		}
		x = x.Next
	}
	return y.Next
}
