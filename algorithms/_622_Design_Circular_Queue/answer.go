package _622_Design_Circular_Queue

type MyCircularQueue struct {
	head, tail int
	data       []int
}

func (this *MyCircularQueue) Len() int {
	if this.tail < this.head {
		return len(this.data) - (this.head - this.tail)
	} else {
		return this.tail - this.head
	}
}

func Constructor(k int) MyCircularQueue {
	q := MyCircularQueue{
		data: make([]int, k+1),
	}
	return q
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.Len() == len(this.data)-1 {
		return false
	}
	this.data[this.tail] = value
	this.tail = (this.tail + 1) % len(this.data)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.Len() == 0 {
		return false
	}
	this.head = (this.head + 1) % len(this.data)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.head+this.Len()-1)%len(this.data)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Len() == len(this.data)-1
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
