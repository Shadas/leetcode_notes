package _2095_Delete_the_Middle_Node_of_a_Linked_List

import "github.com/shadas/leetcode_notes/utils/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteMiddle(head *linkedlist.IntListNode) *linkedlist.IntListNode {
	return deleteMiddleFastSlow(head)
}

func deleteMiddleFastSlow(head *linkedlist.IntListNode) *linkedlist.IntListNode {
	newhead := &linkedlist.IntListNode{Next: head}
	fast, slow := newhead, newhead
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	// rm slow.next
	slow.Next = slow.Next.Next
	return newhead.Next
}
