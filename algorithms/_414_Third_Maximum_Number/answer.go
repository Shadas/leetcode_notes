package _414_Third_Maximum_Number

func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var (
		max         int = nums[0]
		heap            = []int{nums[0]}
		min, minidx     = nums[0], 0
	)
	for _, num := range nums {
		// 随时找最大的
		if num > max {
			max = num
		}
		// 看是否重复
		var repeat bool
		for _, h := range heap {
			if h == num {
				repeat = true
				break
			}
		}
		if !repeat {
			// 处理堆
			if len(heap) < 3 { // 直接插入
				heap = append(heap, num)
				if num < min {
					min = num
					minidx = len(heap) - 1
				}
			} else {
				if num > min { // 先删再插
					if minidx == 2 {
						heap = heap[0:minidx]
					} else {
						heap = append(heap[0:minidx], heap[minidx+1:]...)
					}
					heap = append(heap, num)
					// 找最小的
					min, minidx = heap[0], 0
					for idx, h := range heap {
						if h < min {
							min = h
							minidx = idx
						}
					}
				}
			}
		}
	}
	if len(heap) < 3 {
		return max
	}

	return min

}
