package _143_Reorder_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	reorderListWithStack(head)
}

func reorderListWithStack(head *ListNode) {
	var (
		s    = []int{}
		tmp  = head
		flag bool
	)
	if head == nil {
		return
	}
	for tmp != nil {
		s = append(s, tmp.Val)
		tmp = tmp.Next
	}
	tmp = &ListNode{
		Next: head,
	}
	for len(s) > 0 {
		var value int
		if !flag { // 取头
			value = s[0]
			s = s[1:]
		} else { // 取尾
			value = s[len(s)-1]
			s = s[:len(s)-1]
		}
		tmp.Next.Val = value
		tmp = tmp.Next
		flag = !flag // 换向
	}
	head = head.Next
	return
}
