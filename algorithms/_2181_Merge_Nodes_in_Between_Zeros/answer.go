package _2181_Merge_Nodes_in_Between_Zeros

import "github.com/shadas/leetcode_notes/utils/linkedlist"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *linkedlist.IntListNode) *linkedlist.IntListNode {
	return mergeNodesSliceRun(head)
}

func mergeNodesSliceRun(head *linkedlist.IntListNode) *linkedlist.IntListNode {
	var (
		tmpSum int = -1
		s      []int
	)
	for head != nil {
		if head.Val == 0 {
			if tmpSum != -1 {
				s = append(s, tmpSum)
			}
			tmpSum = 0
		}
		tmpSum += head.Val
		head = head.Next
	}
	x := &linkedlist.IntListNode{}
	y := x
	for _, i := range s {
		x.Next = &linkedlist.IntListNode{Val: i}
		x = x.Next
	}
	return y.Next
}
