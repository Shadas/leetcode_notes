package _24_Swap_Nodes_in_Pairs

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

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ret := head
	count := 0
	var front, back *ListNode
	for head != nil {
		count++
		if count%2 != 0 {
			front = head
		} else {
			back = head
			front.Val, back.Val = back.Val, front.Val
		}
		head = head.Next
	}
	return ret
}
