package _641_Design_Circular_Deque

type MyCircularDeque struct {
	data       []int
	head, tail int
}

func Constructor(k int) MyCircularDeque {
	q := MyCircularDeque{data: make([]int, k+1)}
	return q
}

func (this *MyCircularDeque) Len() int {
	if this.tail > this.head {
		return this.tail - this.head
	} else {
		return len(this.data) - (this.head - this.tail)
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head = (this.head + len(this.data) - 1) % len(this.data)
	this.data[this.head] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.tail] = value
	this.tail = (this.tail + 1) % len(this.data)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.data)
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail - 1 + len(this.data)) % len(this.data)
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.tail-1+len(this.data))%len(this.data)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Len() == len(this.data)-1
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
