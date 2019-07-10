package _460_LFU_Cache

import "fmt"

type Pair struct {
	Value int
	Freq  int
}

type LFUCache struct {
	MinFreq  int
	Capacity int
	Mp       map[int]Pair
	Mf       map[int][]int
}

func Constructor(capacity int) LFUCache {
	lc := LFUCache{
		MinFreq:  0,
		Capacity: capacity,
		Mp:       make(map[int]Pair),
		Mf:       make(map[int][]int),
	}
	return lc
}

func (this *LFUCache) Get(key int) int {
	var (
		p   Pair
		ok  bool
		ret int
	)
	defer func() {
		fmt.Printf("GET %d->%d: %+v\n\n", key, ret, this)
	}()
	// 如果检索不到记录，返回-1
	if p, ok = this.Mp[key]; !ok {
		ret = -1
		return ret
	}
	// 有检索记录
	var (
		oldFreq = p.Freq // 旧频率
		newFreq = oldFreq + 1
		ofl     = this.Mf[oldFreq] // 旧频率列表
		ofln    = []int{}
		nfl     []int
	)
	ret = p.Value // 获得返回值
	// 更新旧的频率列表
	for _, k := range ofl {
		if k != key {
			ofln = append(ofln, k)
		}
	}
	this.Mf[oldFreq] = ofln
	// 更新mp
	p.Freq = newFreq
	this.Mp[key] = p
	// 更新新的mf
	if nfl, ok = this.Mf[newFreq]; !ok {
		nfl = []int{}
	}
	nfl = append([]int{key}, nfl...)
	this.Mf[newFreq] = nfl
	// 更新最低频率标记
	if this.MinFreq == oldFreq { // 如果删除的项的频率对应最低频率
		if len(ofln) == 0 {
			this.MinFreq = newFreq
		}
	}

	return ret
}

func (this *LFUCache) Put(key int, value int) {
	defer fmt.Printf("PUT %d-%d: %+v\n\n", key, value, this)
	// 如果之前有这一项，则更新值，同时把其在频率行中提前
	if p, ok := this.Mp[key]; ok {
		p.Value = value
		this.Mp[key] = p
		fl := this.Mf[p.Freq]
		nl := []int{key}
		for _, x := range fl {
			if x != key {
				nl = append(nl, x)
			}
		}
		this.Mf[p.Freq] = nl
		return
	}
	// 如果没有,则需要增加，增加要看容量是否已经满了
	if len(this.Mp) < this.Capacity { // 如果未满，直接插入即可
		this.Mp[key] = Pair{
			Value: value,
			Freq:  1,
		}
		if f, ok := this.Mf[1]; ok {
			f = append([]int{key}, f...)
			this.Mf[1] = f
		} else {
			this.Mf[1] = []int{key}
		}
		this.MinFreq = 1
		return
	}
	// 如果需要增加缓存项，且缓存项已满，则需要删除
	// 先删除
	// 找到删除的key
	fl := this.Mf[this.MinFreq]
	if len(fl) == 0 {
		return
	}
	delKey := fl[len(fl)-1]
	// 删除记录
	delete(this.Mp, delKey)
	// 删除并更新频率记录
	fl = fl[0 : len(fl)-1]
	this.Mf[this.MinFreq] = fl
	// 再添加项目
	this.Mp[key] = Pair{
		Value: value,
		Freq:  1,
	}
	if f, ok := this.Mf[1]; ok {
		f = append([]int{key}, f...)
		this.Mf[1] = f
	} else {
		this.Mf[1] = []int{key}
	}
	this.MinFreq = 1
	return
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
