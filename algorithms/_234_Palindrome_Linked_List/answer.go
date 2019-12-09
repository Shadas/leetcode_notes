package _234_Palindrome_Linked_List

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

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var (
		fast, slow   = head.Next, head
		cur, newHead *ListNode
	)
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	cur = slow.Next
	slow.Next = nil
	newHead = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = tmp
	}
	for head != nil && newHead != nil {
		if head.Val != newHead.Val {
			return false
		}
		head = head.Next
		newHead = newHead.Next
	}

	return true
}
