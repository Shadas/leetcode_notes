package _142_Linked_List_Cycle_2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	return detectCycleWithPointer(head)
}

func detectCycleWithPointer(head *ListNode) *ListNode {
	var (
		hhead = &ListNode{
			Next: head,
		}
		fast, slow = hhead, hhead
	)

	for {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	slow = hhead
	for {
		if slow == fast {
			return slow
		}
		slow = slow.Next
		fast = fast.Next
	}
	return nil

}
