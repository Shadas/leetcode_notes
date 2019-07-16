package _155_Min_Stack

type MinStack struct {
	l []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		l: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.l = append(this.l, x)
}

func (this *MinStack) Pop() {
	if len(this.l) > 0 {
		this.l = this.l[:len(this.l)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.l) > 0 {
		return this.l[len(this.l)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.l) == 0 {
		return 0
	}
	var ret = this.l[0]
	for _, v := range this.l {
		if v < ret {
			ret = v
		}
	}
	return ret
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
