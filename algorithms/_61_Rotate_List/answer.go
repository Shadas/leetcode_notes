package _61_Rotate_List

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

func rotateRight(head *ListNode, k int) *ListNode {
	var (
		valueList    = []int{}
		newValueList []int
		offset       int
		tmpHead      *ListNode
		ret          *ListNode
	)
	if head == nil {
		return nil
	}
	tmpHead = head
	for {
		valueList = append(valueList, tmpHead.Val)
		if tmpHead.Next == nil {
			break
		} else {
			tmpHead = tmpHead.Next
		}
	}
	offset = k % len(valueList)
	if offset == 0 {
		return head
	}
	newValueList = append(valueList[len(valueList)-offset:], valueList[0:len(valueList)-offset]...)
	ret = &ListNode{}
	tmpHead = ret
	for k, value := range newValueList {
		tmpHead.Val = value
		if k != len(newValueList)-1 {
			tmpHead.Next = &ListNode{}
			tmpHead = tmpHead.Next
		}
	}
	return ret
}
