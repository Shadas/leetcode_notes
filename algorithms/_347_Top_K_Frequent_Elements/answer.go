package _347_Top_K_Frequent_Elements

type Freq struct {
	num  int
	freq int
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var (
		countMap = make(map[int]int)
		heap     = []Freq{}
		minidx   int
		min      Freq
		ret      []int
	)

	for _, num := range nums {
		// 获得计数
		count, _ := countMap[num]
		count++
		countMap[num] = count
		// 处理堆
		if len(heap) < k { // 直接插入
			tmp := Freq{
				num:  num,
				freq: count,
			}
			var exist bool
			var existIdx int
			for idx, h := range heap {
				if h.num == num {
					exist = true
					existIdx = idx
					break
				}
			}
			if exist {
				h := heap[existIdx]
				if count > h.freq {
					h.freq = count
				}
				heap[existIdx] = h
			} else {
				heap = append(heap, tmp)
			}

			if len(heap) == 1 { // 初始化情形
				minidx = 0
			} else {
				if min.freq > tmp.freq {
					if exist {
						minidx = existIdx
					} else {
						minidx = len(heap) - 1
					}
				}
			}
			min = heap[minidx]
		} else {
			// 找有没有重复
			var exist bool
			var existIdx int
			for idx, h := range heap {
				if h.num == num {
					exist = true
					existIdx = idx
				}
			}
			if exist {
				h := heap[existIdx]
				h.freq = count
				heap[existIdx] = h
			} else {
				if count > min.freq { // 需要删除
					if minidx == k-1 {
						heap = heap[0 : k-1]
					} else {
						heap = append(heap[0:minidx], heap[minidx+1:]...)
					}
					heap = append(heap, Freq{
						num:  num,
						freq: count,
					})

				}
				// 找新的最小的
				minidx = 0
				min = heap[0]
				for idx, h := range heap {
					if h.freq < min.freq {
						min = h
						minidx = idx
					}
				}
			}

		}
	}
	for _, h := range heap {
		ret = append(ret, h.num)
	}

	return ret

}
