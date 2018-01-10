package _2_Add_Two_Numbers

import (
	"fmt"
	"strings"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
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

	iswrong := false

	want := []int{7, 0, 8}

	for i := 0; i < 3; i++ {
		if ret.Val == want[i] {
			if i == 2 {
				if ret.Next == nil {
					break
				}
			} else {
				if ret.Next != nil {
					ret = ret.Next
				} else {
					iswrong = true
					break
				}
			}
		} else {
			iswrong = true
			break
		}
	}

	var rets []string
	for {
		rets = append(rets, fmt.Sprintf("%v", ret.Val))
		if ret.Next != nil {
			ret = ret.Next
		} else {
			break
		}
	}
	retstr := strings.Join(rets, ",")

	if iswrong {
		t.Errorf("Your answer is %v, it should be %v", "["+retstr+"]", want)
	}

}
