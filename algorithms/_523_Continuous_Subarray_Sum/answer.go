package _523_Continuous_Subarray_Sum

func checkSubarraySum(nums []int, k int) bool {
	return checkSubarraySumWithSums(nums, k)
}

func checkSubarraySumWithSums(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	var (
		sum          int
		remainderMap = map[int]int{0: -1} // map[remainder]idx
	)
	for idx, num := range nums {
		sum += num
		rmd := sum % k
		if tidx, ok := remainderMap[rmd]; ok {
			if idx-tidx <= 1 { // 连续2个以上才有效
				continue
			}
			return true
		}
		remainderMap[rmd] = idx
	}
	return false
}
