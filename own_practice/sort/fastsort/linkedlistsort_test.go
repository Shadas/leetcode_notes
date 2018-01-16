package fastsort

import (
	"testing"

	. "leetcode_notes/utils/linkedlist"
)

func TestLinkedListFastSort1(t *testing.T) {
	l1 := GenerateIntLinkedList([]int{3, 1, 2})
	ret := LinkedListFastSort1(l1)
	retl1 := GenerateIntLinkedList([]int{1, 2, 3})
	if !IsTwoIntLinkedListEqual(l1, retl1) {
		t.Error("error test1 with", IntLinkedList2slice(ret))
	}
}
