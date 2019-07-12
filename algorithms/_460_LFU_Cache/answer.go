package _460_LFU_Cache

import (
	"fmt"
	"strconv"
	"strings"
)

// 测试时可调用 transCall([]string{"LFUCache", "put", "get"...}, transArg("[[2],[1,1],[1]]...")) 类似的方式进行测试，方便调试
// 测试用，处理arg参数
func transArg(str string) (ret [][]int) {
	str = strings.TrimSpace(strings.TrimRight(strings.TrimLeft(str, "["), "]"))
	args := strings.Split(str, "],[")
	for _, arg := range args {
		as := strings.Split(strings.TrimSpace(arg), ",")
		tmpAs := []int{}
		for _, a := range as {
			i, _ := strconv.Atoi(a)
			tmpAs = append(tmpAs, i)
		}
		ret = append(ret, tmpAs)
	}
	return
}

// 测试用，处理调用
func transCall(fn []string, arg [][]int) {
	if len(fn) != len(arg) {
		fmt.Println("err input")
		return
	}
	var obj LFUCache
	for i := 0; i < len(fn); i++ {
		switch fn[i] {
		case "LFUCache":
			obj = Constructor(arg[i][0])
		case "put":
			obj.Put(arg[i][0], arg[i][1])
		case "get":
			obj.Get(arg[i][0])
		}

	}

}

type Pair struct {
	Value int
	Freq  int
}

type LFUCache struct {
	MinFreq  int
	Capacity int
	Mp       map[int]Pair
	Mf       map[int]*[]int
}

func Constructor(capacity int) LFUCache {
	lc := LFUCache{
		MinFreq:  0,
		Capacity: capacity,
		Mp:       make(map[int]Pair),
		Mf:       make(map[int]*[]int),
	}
	return lc
}

// 打印内容仅限于调试，跑的时候一定要去掉，不然时间会超限。
func (this *LFUCache) Get(key int) int {
	var (
		p   Pair
		ok  bool
		ret int
	)
	// defer func() {
	// 	fmt.Printf("GET %d->%d: %+v\n\n", key, ret, this)
	// }()
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
		nfl     *[]int
		delIdx  int
	)
	ret = p.Value // 获得返回值
	if len(*ofl) == 0 {
		return -1
	}
	// 更新旧的频率列表
	for i, k := range *ofl {
		if k == key {
			delIdx = i
			break
		}
	}
	if delIdx == len(*ofl)-1 {
		*ofl = append((*ofl)[:delIdx])
	} else {
		*ofl = append((*ofl)[:delIdx], (*ofl)[delIdx+1:]...)
	}
	this.Mf[oldFreq] = ofl
	// 更新mp
	p.Freq = newFreq
	this.Mp[key] = p
	// 更新新的mf
	if nfl, ok = this.Mf[newFreq]; !ok {
		nfl = &[]int{}
	}
	*nfl = append(*nfl, key)
	this.Mf[newFreq] = nfl
	// 更新最低频率标记
	if this.MinFreq == oldFreq { // 如果删除的项的频率对应最低频率
		if len(*ofl) == 0 {
			this.MinFreq = newFreq
		}
	}

	return ret
}

func (this *LFUCache) Put(key int, value int) {
	// defer fmt.Printf("PUT %d-%d: %+v\n\n", key, value, this)
	// 如果之前有这一项，则更新值，同时把其在频率行中提前
	if p, ok := this.Mp[key]; ok {
		var (
			delIdx int
			of     = p.Freq
			nf     = of + 1
			ofl    = this.Mf[of]
		)
		if ofl == nil || len(*ofl) == 0 {
			return
		}
		p.Value = value
		p.Freq = nf
		this.Mp[key] = p
		// 更新旧的频率列表
		for i, k := range *ofl {
			if k == key {
				delIdx = i
				break
			}
		}
		if delIdx == len(*ofl)-1 {
			*ofl = append((*ofl)[:delIdx])
		} else {
			*ofl = append((*ofl)[:delIdx], (*ofl)[delIdx+1:]...)
		}
		this.Mf[of] = ofl

		var nfl *[]int
		if nfl, ok = this.Mf[nf]; !ok {
			nfl = &[]int{}
		}
		*nfl = append(*nfl, key)
		this.Mf[nf] = nfl

		if this.MinFreq == of && len(*this.Mf[of]) == 0 {
			this.MinFreq = nf
		}
		return
	}
	// 如果没有,则需要增加，增加要看容量是否已经满了
	if len(this.Mp) < this.Capacity { // 如果未满，直接插入即可
		this.Mp[key] = Pair{
			Value: value,
			Freq:  1,
		}
		var (
			f  *[]int
			ok bool
		)
		if f, ok = this.Mf[1]; !ok {
			f = &[]int{}
		}
		*f = append(*f, key)
		this.Mf[1] = f
		this.MinFreq = 1
		return
	}
	// 如果需要增加缓存项，且缓存项已满，则需要删除
	// 先删除
	// 找到删除的key
	fl := this.Mf[this.MinFreq]
	if fl == nil || len(*fl) == 0 {
		return
	}
	delKey := (*fl)[0]
	// 删除记录
	delete(this.Mp, delKey)
	// 删除并更新频率记录
	*fl = (*fl)[1:]
	this.Mf[this.MinFreq] = fl
	// 再添加项目
	this.Mp[key] = Pair{
		Value: value,
		Freq:  1,
	}
	if f, ok := this.Mf[1]; ok {
		*f = append((*f), key)
		this.Mf[1] = f
	} else {
		this.Mf[1] = &[]int{key}
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
