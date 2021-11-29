package _382_Linked_List_Random_Node

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	var (
		scope       = 1
		chosenValue = 0
	)
	curr := this.head
	for curr != nil {
		if rand.Float64() < 1.0/float64(scope) {
			chosenValue = curr.Val
		}
		scope += 1
		curr = curr.Next
	}
	return chosenValue
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
