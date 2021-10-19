package _232_Implement_Queue_using_Stacks

import "sync"

type MyQueue struct {
	in, out []int // 两个栈，分别处理进出
	top     int

	rwm sync.RWMutex
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	q := MyQueue{}
	return q
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.rwm.Lock()
	defer this.rwm.Unlock()
	if len(this.in) == 0 {
		this.top = x
	}
	this.in = append(this.in, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.rwm.Lock()
	defer this.rwm.Unlock()
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			// pop
			x := this.in[len(this.in)-1]
			this.in = this.in[:len(this.in)-1]
			// push
			this.out = append(this.out, x)
		}
	}
	ret := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return ret
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.rwm.RLock()
	defer this.rwm.RUnlock()
	if len(this.out) == 0 {
		return this.top
	}
	return this.out[len(this.out)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	this.rwm.RLock()
	defer this.rwm.RUnlock()
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
