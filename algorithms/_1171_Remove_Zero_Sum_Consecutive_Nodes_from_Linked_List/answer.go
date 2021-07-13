package _1171_Remove_Zero_Sum_Consecutive_Nodes_from_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	sum, p, m := 0, head, make(map[int]*ListNode)
	for p != nil {
		sum += p.Val
		if sum == 0 {
			head = p.Next
			p = head
			m = make(map[int]*ListNode)
		} else {
			if node, ok := m[sum]; ok {
				node.Next = p.Next
				p = head
				sum, m = 0, make(map[int]*ListNode)
			} else {
				m[sum] = p
				p = p.Next
			}
		}
	}
	return head
}
