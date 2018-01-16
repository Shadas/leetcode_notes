package fastsort

type ListNode struct {
	Val  int
	Next *ListNode
}

//
func LinkedListFastSort1(l *ListNode) *ListNode {
	head := l
	linkedListFastSort1(head, nil)
	return head
}

func linkedListFastSort1(head, end *ListNode) {
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
