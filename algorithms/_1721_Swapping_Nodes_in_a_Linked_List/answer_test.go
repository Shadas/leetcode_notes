package _1721_Swapping_Nodes_in_a_Linked_List

import (
	"fmt"
	"testing"
)

func TestSwapNodes(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	k := 2
	ret := swapNodes(head, k)
	fmt.Println(ret)
}
