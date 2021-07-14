package _25_Reverse_Nodes_in_k_Group

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

func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		count int
		tmp   = &ListNode{}
		ret   = &ListNode{}
		tret  = ret         // ret的尾节点，用于接下一串链接
		ttail = &ListNode{} // 子串的尾节点，用于更新ret的尾结点
	)
	for head != nil {
		if count < k {
			x := &ListNode{
				Val:  head.Val,
				Next: tmp.Next,
			}
			tmp.Next = x
			if count == 0 {
				ttail = x
			}
			head = head.Next
			count++
		} else {
			tret.Next = tmp.Next
			tmp = &ListNode{}
			tret = ttail
			count = 0
		}
	}
	if count == k {
		// 把最后的tmp缀上来
		tret.Next = tmp.Next
	} else if count > 0 {
		// 把最后的tmp要反回来
		var (
			tmpTail = &ListNode{}
		)
		for tmp.Next != nil {
			x := &ListNode{
				Val:  tmp.Next.Val,
				Next: tmpTail.Next,
			}
			tmpTail.Next = x
			tmp = tmp.Next
		}
		tret.Next = tmpTail.Next
	}
	return ret.Next
}
