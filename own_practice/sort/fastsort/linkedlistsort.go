package fastsort

import (
	. "leetcode_notes/utils/linkedlist"
)

//
func LinkedListFastSort1(l *IntListNode) *IntListNode {
	head := l
	linkedListFastSort1(head, nil)
	return head
}

func linkedListFastSort1(head, end *IntListNode) {
	if head == nil || head == end {
		return
	}
	p := head.Next // pointer for run
	small := head
	for p != end {
		if p.Val < head.Val {
			small = small.Next
			small.Val, p.Val = p.Val, small.Val
		}
		p = p.Next
	}
	head.Val, small.Val = small.Val, head.Val
	linkedListFastSort1(head, small)
	linkedListFastSort1(small.Next, end)
}
