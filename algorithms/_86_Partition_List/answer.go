package _86_Partition_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	sh, bh := &ListNode{}, &ListNode{}
	sht, bht := sh, bh
	for {
		if head == nil {
			break
		}
		if head.Val < x {
			sht.Next = &ListNode{
				Val: head.Val,
			}
			sht = sht.Next
		} else {
			bht.Next = &ListNode{
				Val: head.Val,
			}
			bht = bht.Next
		}
		head = head.Next
	}
	sht.Next = bh.Next
	return sh.Next

}
