package _23_Merge_K_Sorted_Lists

import (
	"fmt"
	"testing"
)

// GenerateLinkedList generates an linked list
func GenerateLinkedList(vals []int) *ListNode {
	var l *ListNode
	var head *ListNode
	for _, val := range vals {
		if l == nil {
			l = &ListNode{
				Val: val,
			}
			head = l
		} else {
			l.Next = &ListNode{
				Val: val,
			}
			l = l.Next
		}
	}
	return head
}

func PrintListNode(l *ListNode) {
	var s string
	for l != nil {
		s += fmt.Sprintf("%d", l.Val)
		l = l.Next
	}
	fmt.Println(s)
}

func TestMergeKLists(t *testing.T) {
	l1 := GenerateLinkedList([]int{1, 4, 5})
	l2 := GenerateLinkedList([]int{1, 3, 4})
	l3 := GenerateLinkedList([]int{2, 6})
	lists1 := []*ListNode{l1, l2, l3}
	ret := mergeKLists(lists1)
	PrintListNode(ret)

	l4 := GenerateLinkedList([]int{-2, -1, -1, -1})
	l5 := GenerateLinkedList([]int{})
	lists2 := []*ListNode{l4, l5}
	ret = mergeKLists(lists2)
	PrintListNode(ret)
}
