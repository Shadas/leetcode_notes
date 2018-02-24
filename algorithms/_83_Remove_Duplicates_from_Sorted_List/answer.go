package _83_Remove_Duplicates_from_Sorted_List

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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ret := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			if head.Next.Next != nil {
				head.Next = head.Next.Next
			} else {
				head.Next = nil
				break
			}
		} else {
			head = head.Next
		}
	}
	return ret
}
