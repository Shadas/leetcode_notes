package _225_Implement_Stack_using_Queues

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	tmp := []int{}
	for len(this.queue) > 1 {
		x := this.queue[0]
		this.queue = this.queue[1:]
		tmp = append(tmp, x)
	}
	ret := this.queue[0]
	this.queue = tmp
	return ret
}

func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
