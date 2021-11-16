package _560_Subarray_Sum_Equals_K

func subarraySum(nums []int, k int) int {
	return subarraySumWithMap(nums, k)
}

func subarraySumWithMap(nums []int, k int) int {
	var (
		m     = map[int]int{0: 1} // map[sum]count; 初始为0，计一次
		sum   int
		count int
	)
	for _, num := range nums {
		sum += num
		key := sum - k // 去之前的sum里找找有没有和当前sum差k的情况
		if tcount, ok := m[key]; ok {
			count += tcount
		}
		m[sum] += 1
	}
	return count
}
