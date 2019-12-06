package _1_Two_Sum

func twoSum(nums []int, target int) []int {
	// return twoSumForce(nums, target)
	return twoSumHash(nums, target)
}

func twoSumHash(nums []int, target int) []int {
	var (
		m = make(map[int]int)
	)
	for idx, n := range nums {
		m[n] = idx
	}
	for idx1, n := range nums {
		if idx2, ok := m[target-n]; ok && idx1 != idx2 {
			return []int{idx1, idx2}
		}
	}
	return []int{}
}

func twoSumForce(nums []int, target int) []int {
	ret := []int{}
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if k1 >= k2 {
				continue
			}
			if v1+v2 == target {
				ret = append(ret, k1, k2)
				return ret
			}
		}
	}
	return ret
}
