package _725_Split_Linked_List_in_Parts

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

func splitListToParts(root *ListNode, k int) []*ListNode {
	count := 0
	tmp := root
	for tmp != nil {
		count++
		tmp = tmp.Next
	}
	per := count / k
	last := count % k
	ret := []*ListNode{}
	head := root

	for i := 0; i < k; i++ {
		th := &ListNode{}
		thp := th
		for j := 0; j < per; j++ {
			x := &ListNode{Val: head.Val}
			thp.Next = x
			thp = thp.Next
			head = head.Next
		}
		if i < last {
			x := &ListNode{Val: head.Val}
			thp.Next = x
			thp = thp.Next
			head = head.Next
		}
		ret = append(ret, th.Next)
	}
	return ret
}
