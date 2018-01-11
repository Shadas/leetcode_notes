package _21_Merge_Two_Sorted_Lists

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	header := ListNode{}
	ret := &header
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val > l2.Val {
			ret.Next = l2
			l2 = l2.Next
		} else {
			ret.Next = l1
			l1 = l1.Next
		}
		ret = ret.Next
	}
	if l1 == nil {
		ret.Next = l2
	} else {
		ret.Next = l1
	}
	return header.Next
}
