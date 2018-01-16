package fastsort

import (
	"fmt"
	"testing"
)

func makeLinkedList(vals []int) *ListNode {
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

func linkedList2slice(l *ListNode) []int {
	ret := []int{}
	if l == nil {
		return ret
	}
	for {
		ret = append(ret, l.Val)
		if l.Next != nil {
			l = l.Next
		} else {
			break
		}
	}
	return ret
}

func isEqualTwoLinkedList(l1, l2 *ListNode) bool {
	sl1 := linkedList2slice(l1)
	sl2 := linkedList2slice(l2)
	if len(sl1) != len(sl2) {
		return false
	}
	for i := 0; i < len(sl1); i++ {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}
