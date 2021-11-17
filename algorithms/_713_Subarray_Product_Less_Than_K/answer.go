package _713_Subarray_Product_Less_Than_K

func numSubarrayProductLessThanK(nums []int, k int) int {
	return numSubarrayProductLessThanKWithSlidingWindow(nums, k)
}

func numSubarrayProductLessThanKWithSlidingWindow(nums []int, k int) int {
	var (
		count int
		p, q  int     // 前后idx
		prod  int = 1 // 乘积
	)
	for q = 0; q < len(nums); q++ { // 以右界增加处理
		prod *= nums[q]
		for prod >= k && p <= q {
			prod /= nums[p]
			p++
		}
		count += (q - p + 1) // 每向右移动一次右界，增加的连续个数
	}
	return count
}
