package _82_Remove_Duplicates_from_Sorted_List_2

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
	var ret *ListNode
	var retTail *ListNode
	var lastValue int
	var isRepeated bool
	if head.Next == nil {
		return head
	} else {
		lastValue = head.Val
	}
	for head.Next != nil {
		head = head.Next
		if head.Val == lastValue {
			isRepeated = true
		} else {
			if !isRepeated {
				if ret == nil {
					ret = &ListNode{
						Val: lastValue,
					}
					retTail = ret
				} else {
					retTail.Next = &ListNode{
						Val: lastValue,
					}
					retTail = retTail.Next
				}
			} else {
				isRepeated = false
			}
			lastValue = head.Val
		}
	}
	if !isRepeated {
		if ret == nil {
			ret = &ListNode{
				Val: lastValue,
			}
			retTail = ret
		} else {
			retTail.Next = &ListNode{
				Val: lastValue,
			}
			retTail = retTail.Next
		}
	}
	return ret
}
