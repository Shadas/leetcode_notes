package _2_Add_Two_Numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{}
	head := prev
	carry := 0
	for {
		cur := &ListNode{}
		l1val, l2val := 0, 0
		if l1 != nil {
			l1val = l1.Val
			l1 = l1.Next
		} else {
			l1 = nil
		}
		if l2 != nil {
			l2val = l2.Val
			l2 = l2.Next
		} else {
			l2 = nil
		}
		sum := l1val + l2val + carry
		cur.Val = sum % 10
		carry = sum / 10
		prev.Next = cur
		prev = cur

		if l1 == nil && l2 == nil && carry == 0 {
			break
		}

	}

	return head.Next

}
