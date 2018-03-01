package _92_Reverse_Linked_List_2

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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	ret := &ListNode{}
	ret.Next = head
	pre := ret
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	start := pre.Next
	then := start.Next

	for i := 0; i < n-m; i++ {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = start.Next
	}

	return ret.Next
}
