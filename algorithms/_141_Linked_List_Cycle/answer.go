package _141_Linked_List_Cycle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var fast, slow = head, head
	fast = fast.Next.Next
	slow = slow.Next
	for {
		if fast == nil || fast.Next == nil {
			return false
		} else if fast == slow || fast.Next == slow {
			return true
		} else {
			fast = fast.Next.Next
			slow = slow.Next
		}
	}
	return false
}
