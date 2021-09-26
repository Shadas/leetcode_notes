package _215_Kth_Largest_Element_in_an_Array

import "sort"

func findKthLargest(nums []int, k int) int {
	//return findKthLargestWithKLengthArray(nums, k)
	return findKthLargestAfterSortingAllElements(nums, k)
}

// 全排序然后取第K位
func findKthLargestAfterSortingAllElements(nums []int, k int) int {
	sort.Ints(nums) // 升序
	return nums[len(nums)-k]
}

// 其实没有用到堆，只是用了一个数组，每次添加都遍历一遍找最小，其实效率并不好
func findKthLargestWithKLengthArray(nums []int, k int) int {
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
