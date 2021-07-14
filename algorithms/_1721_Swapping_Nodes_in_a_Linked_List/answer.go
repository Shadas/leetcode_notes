package _1721_Swapping_Nodes_in_a_Linked_List

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var s string
	for l != nil {
		s += fmt.Sprintf("%d", l.Val)
		l = l.Next
	}
	return s
}

func swapNodes(head *ListNode, k int) *ListNode {
	ls, lf := head, head
	for i := 0; i < k-1; i++ {
		lf = lf.Next
	}
	fn := lf
	for lf.Next != nil {
		ls = ls.Next
		lf = lf.Next
	}
	ln := ls
	fn.Val, ln.Val = ln.Val, fn.Val
	return head
}
