package _725_Split_Linked_List_in_Parts

import (
	"fmt"
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	k := 5
	ret := splitListToParts(l, k)
	for _, tl := range ret {
		fmt.Println(tl)
	}
}
