package _876_Middle_of_the_Linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var (
		fast, slow *ListNode = head, head
		slowMove   bool      = true
	)
	if head == nil {
		return nil
	}
	for {
		if fast.Next == nil {
			break
		}
		if slowMove {
			slow = slow.Next
		}
		slowMove = !slowMove
		fast = fast.Next
	}
	return slow
}
