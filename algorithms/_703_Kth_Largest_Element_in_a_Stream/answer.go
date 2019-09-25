package _703_Kth_Largest_Element_in_a_Stream

type KthLargest struct {
	queue []int
	k     int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		k: k,
	}
	for _, n := range nums {
		if len(kl.queue) == 0 {
			kl.queue = append(kl.queue, n)
		} else {
			if n <= kl.queue[len(kl.queue)-1] {
				kl.queue = append(kl.queue, n)
			} else if n > kl.queue[0] {
				kl.queue = append([]int{n}, kl.queue...)
			} else {
				for i := len(kl.queue) - 2; i >= 0; i-- {
					if n <= kl.queue[i] {
						kl.queue = append(kl.queue[0:i+1], append([]int{n}, kl.queue[i+1:]...)...)
						break
					}
				}
			}
		}
	}
	// 取前k个
	if len(kl.queue) >= k {
		kl.queue = kl.queue[0:k]
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	if len(this.queue) == 0 {
		this.queue = append(this.queue, val)
	} else {
		if val <= this.queue[len(this.queue)-1] {
			this.queue = append(this.queue, val)
		} else if val > this.queue[0] {
			this.queue = append([]int{val}, this.queue...)
		} else {
			for i := len(this.queue) - 2; i >= 0; i-- {
				if val <= this.queue[i] {
					this.queue = append(this.queue[0:i+1], append([]int{val}, this.queue[i+1:]...)...)
					break
				}
			}
		}
	}
	var ret int
	if len(this.queue) >= this.k {
		this.queue = this.queue[0:this.k]
		ret = this.queue[this.k-1]
	} else {
		ret = -1
	}
	return ret
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
