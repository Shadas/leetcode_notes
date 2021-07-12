package _328_Odd_Even_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	h1, h2 := &ListNode{}, &ListNode{}
	l1, l2 := h1, h2
	var isEven bool
	for head != nil {
		if !isEven {
			l1.Next = &ListNode{
				Val: head.Val,
			}
			l1 = l1.Next
		} else {
			l2.Next = &ListNode{
				Val: head.Val,
			}
			l2 = l2.Next
		}
		head = head.Next
		isEven = !isEven
	}
	l1.Next = h2.Next
	return h1.Next
}
