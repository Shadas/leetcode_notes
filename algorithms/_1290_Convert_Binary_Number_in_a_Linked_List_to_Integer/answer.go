package _1290_Convert_Binary_Number_in_a_Linked_List_to_Integer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var val int
	for {
		if head == nil {
			break
		}
		val <<= 1
		val += head.Val
		head = head.Next
	}
	return val
}
