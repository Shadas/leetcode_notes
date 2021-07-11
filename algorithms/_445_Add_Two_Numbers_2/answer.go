package _445_Add_Two_Numbers_2

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var s string
	for l != nil {
		s += fmt.Sprintf("%d", l.Val)
		l = l.Next
	}
	return s
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		s1, s2 []int
		ret    = &ListNode{}
	)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	var addMore bool
	for len(s1) > 0 && len(s2) > 0 {
		t1 := s1[len(s1)-1]
		t2 := s2[len(s2)-1]
		s1 = s1[:len(s1)-1]
		s2 = s2[:len(s2)-1]
		sum := t1 + t2
		addMore = dealSum(sum, ret, addMore)
	}
	// 等长部分处理完
	if len(s2) > 0 {
		s1 = s2 // 统一到s1处理
	}
	// 处理剩余
	for len(s1) > 0 {
		t1 := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		sum := t1
		addMore = dealSum(sum, ret, addMore)
	}
	// 处理最后进位
	if addMore {
		tmp := &ListNode{
			Val:  1,
			Next: ret.Next,
		}
		ret.Next = tmp
	}
	return ret.Next
}

func dealSum(sum int, ret *ListNode, addMore bool) (newAddMore bool) {
	if addMore {
		sum += 1
		newAddMore = false
	}
	v := sum % 10
	newAddMore = sum/10 == 1
	if ret.Next == nil { // 说明是最低位
		ret.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
	} else {
		tmp := &ListNode{
			Val:  v,
			Next: ret.Next,
		}
		ret.Next = tmp
	}
	return
}
