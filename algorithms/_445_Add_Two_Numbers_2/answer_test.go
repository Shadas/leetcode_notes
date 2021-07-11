package _445_Add_Two_Numbers_2

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	ret := addTwoNumbers(l1, l2)
	fmt.Println(ret)

	l3 := &ListNode{
		Val:  0,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  0,
		Next: nil,
	}
	ret1 := addTwoNumbers(l3, l4)
	fmt.Println(ret1)
}
