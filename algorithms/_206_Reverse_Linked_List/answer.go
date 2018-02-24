package _206_Reverse_Linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var ret *ListNode
	for head != nil {
		tail := &ListNode{
			Val:  head.Val,
			Next: ret,
		}
		ret = tail
		head = head.Next
	}
	return ret
}
