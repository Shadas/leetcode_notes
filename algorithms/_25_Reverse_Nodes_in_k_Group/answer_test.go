package _25_Reverse_Nodes_in_k_Group

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	var (
		l   *ListNode
		k   int
		ret *ListNode
	)
	l = &ListNode{
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
	k = 3
	ret = reverseKGroup(l, k)
	fmt.Println(ret)

	l = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	k = 2
	ret = reverseKGroup(l, k)
	fmt.Println(ret)
}
