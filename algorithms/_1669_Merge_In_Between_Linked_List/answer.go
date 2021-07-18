package _1669_Merge_In_Between_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	l1, l2 := list1, list2
	for ; a > 1; a-- {
		l1 = l1.Next
		b--
	}
	pre := l1
	for ; b >= 0; b-- {
		l1 = l1.Next
	}
	last := l1
	pre.Next = list2
	for l2.Next != nil {
		l2 = l2.Next
	}
	l2.Next = last
	return list1
}
