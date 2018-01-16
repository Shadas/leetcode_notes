package _148_Sort_List

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

func sortList(head *ListNode) *ListNode {
	fastSortListRecursion(head, nil)
	return head
}

func fastSortListRecursion(head *ListNode, end *ListNode) {
	if head == nil || head == end {
		return
	}
	p := head.Next
	small := head
	for p != end {
		if p.Val < head.Val {
			small = small.Next
			p.Val, small.Val = small.Val, p.Val
		}
		p = p.Next
	}
	head.Val, small.Val = small.Val, head.Val
	fastSortListRecursion(head, small)
	fastSortListRecursion(small.Next, end)
}
