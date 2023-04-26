package _2130_Maximum_Twin_Sum_of_Linked_List

import (
	"math"

	"github.com/shadas/leetcode_notes/utils/linkedlist"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *linkedlist.IntListNode) int {
	return pairSumSlice(head)
}

func pairSumSlice(head *linkedlist.IntListNode) int {
	var x []int
	for head != nil {
		x = append(x, head.Val)
		head = head.Next
	}
	var max = math.MinInt
	for i := 0; i < len(x)/2; i++ {
		tmp := x[i] + x[len(x)-1-i]
		if tmp > max {
			max = tmp
		}
	}
	return max
}
