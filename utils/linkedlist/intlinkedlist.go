package linkedlist

type IntListNode struct {
	Val  int
	Next *IntListNode
}

// IsTwoIntLinkedListEqual compares two int linked list, return if they are the same or not.
func IsTwoIntLinkedListEqual(l1, l2 *IntListNode) bool {
	sl1 := IntLinkedList2slice(l1)
	sl2 := IntLinkedList2slice(l2)
	if len(sl1) != len(sl2) {
		return false
	}
	for i := 0; i < len(sl1); i++ {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}

// IntLinkedList2slice transfer an int linked list to an int slice.
func IntLinkedList2slice(l *IntListNode) []int {
	ret := []int{}
	if l == nil {
		return ret
	}
	for {
		ret = append(ret, l.Val)
		if l.Next != nil {
			l = l.Next
		} else {
			break
		}
	}
	return ret
}

// GenerateIntLinkedList generates an int linked list
func GenerateIntLinkedList(vals []int) *IntListNode {
	var l *IntListNode
	var head *IntListNode
	for _, val := range vals {
		if l == nil {
			l = &IntListNode{
				Val: val,
			}
			head = l
		} else {
			l.Next = &IntListNode{
				Val: val,
			}
			l = l.Next
		}
	}
	return head
}
