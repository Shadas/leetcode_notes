package _933_Number_of_Recent_Calls

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	rc := RecentCounter{}
	return rc
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.queue) != 0 {
		tmp := this.queue[0]
		if t-tmp > 3000 {
			this.queue = this.queue[1:]
		} else {
			break
		}
	}
	this.queue = append(this.queue, t)
	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
