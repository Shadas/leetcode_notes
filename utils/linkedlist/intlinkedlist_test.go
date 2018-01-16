package linkedlist

import (
	"testing"
)

func TestIntLinkedlistTools(t *testing.T) {
	sl := []int{3, 1, 2}
	l := GenerateIntLinkedList(sl)
	if l.Val != 3 || l.Next == nil ||
		l.Next.Val != 1 || l.Next.Next == nil ||
		l.Next.Next.Val != 2 || l.Next.Next.Next != nil {
		t.Error("generate error.")
	}

	l2s := IntLinkedList2slice(l)
	if len(l2s) != len(sl) || l2s[0] != sl[0] ||
		l2s[1] != sl[1] || l2s[2] != sl[2] {
		t.Error("linkedlist to slice error.")
	}

	l1 := GenerateIntLinkedList(sl)
	if isEqual := IsTwoIntLinkedListEqual(l, l1); !isEqual {
		t.Error("is equal error.")
	}
}
