package _239_Sliding_Window_Maximum

import (
	"container/list"
)

func maxSlidingWindow(nums []int, k int) []int {
	return maxSlidingWindowWithQueue(nums, k)
}

// 双端单调队列，为了移动时计算对比窗口长度，队列中不放具体的值，而放下标
func maxSlidingWindowWithQueue(nums []int, k int) []int {
	if k > len(nums) || len(nums) == 0 { // bad case
		return []int{}
	}
	var (
		ret []int
		qs  = list.New()
	)
	// 初始化构造窗口
	for i := 0; i < k; i++ {
		for qs.Len() != 0 && nums[qs.Back().Value.(int)] < nums[i] {
			qs.Remove(qs.Back())
		}
		qs.PushBack(i)
	}
	ret = append(ret, nums[qs.Front().Value.(int)])
	// 开始移动
	for i := k; i < len(nums); i++ {
		if (i - qs.Front().Value.(int) + 1) > k {
			qs.Remove(qs.Front())
		}
		for qs.Len() != 0 && nums[qs.Back().Value.(int)] < nums[i] {
			qs.Remove(qs.Back())
		}
		qs.PushBack(i)
		ret = append(ret, nums[qs.Front().Value.(int)])
	}
	return ret
}
