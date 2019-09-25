package _215_Kth_Largest_Element_in_an_Array

func findKthLargest(nums []int, k int) int {
	var (
		heap        []int
		min, minidx int
	)
	for _, num := range nums {
		if len(heap) == 0 {
			heap = append(heap, num)
			min = num
			minidx = 0
		} else {
			if len(heap) < k { // 直接插入
				heap = append(heap, num)
				if num < min {
					min = num
					minidx = len(heap) - 1
				}
			} else {
				if num > min { // 需要删除
					if minidx == len(heap)-1 {
						heap = heap[0 : len(heap)-1]
					} else {
						heap = append(heap[0:minidx], heap[minidx+1:]...)
					}
					heap = append(heap, num)
					// 重新找最小的
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
	return min
}
