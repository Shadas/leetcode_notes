package _160_Intersection_of_Two_Linked_Lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	return getIntersectionNodeWithLengthWalk(headA, headB)
}

func getIntersectionNodeWithLengthWalk(headA, headB *ListNode) *ListNode {
	var (
		n1 = headA
		n2 = headB
	)
	for n1 != nil && n2 != nil {
		n1 = n1.Next
		n2 = n2.Next
	}
	for n1 != nil { // 说明headA长
		headA = headA.Next
		n1 = n1.Next
	}
	for n2 != nil {
		headB = headB.Next
		n2 = n2.Next
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}
