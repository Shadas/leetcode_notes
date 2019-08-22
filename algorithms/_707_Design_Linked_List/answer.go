package _707_Design_Linked_List

type MyLinkedList struct {
	Length int
	Next   *MyLinkedList
	Val    int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.Length || index < 0 {
		return -1
	}
	tmp := this
	for i := 0; ; i++ {
		if i == index {
			tn := tmp.Next
			return tn.Val
		}
		tmp = tmp.Next
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := &MyLinkedList{
		Val:  val,
		Next: this.Next,
	}
	this.Next = node
	this.Length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	tmp := this
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &MyLinkedList{
		Val: val,
	}
	this.Length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Length {
		return
	}
	if index < 0 && this.Length == 0 {
		index = 0
	}
	tmp := this
	for i := 0; ; i++ {
		if i == index {
			tn := tmp.Next
			node := &MyLinkedList{
				Val: val,
			}
			if tn != nil {
				node.Next = tn
			}
			tmp.Next = node
			this.Length++
			return
		}
		tmp = tmp.Next
	}
	return
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Length || index < 0 {
		return
	}
	tmp := this
	for i := 0; ; i++ {
		if i == index {
			tn := tmp.Next
			tmp.Next = tn.Next
			tn = nil
			this.Length--
			return
		}
		tmp = tmp.Next
	}
	return
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
