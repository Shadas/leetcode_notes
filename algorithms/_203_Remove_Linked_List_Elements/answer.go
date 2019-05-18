package _203_Remove_Linked_List_Elements

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	ret := &ListNode{
		Next: head,
	}
	current := head
	last := ret
	for {
		tmp := current
		if tmp == nil {
			break
		}
		if tmp.Val == val {
			last.Next = tmp.Next
		} else {
			last = tmp
		}
		current = tmp.Next

	}
	return ret.Next
}
