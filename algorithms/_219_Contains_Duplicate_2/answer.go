package _219_Contains_Duplicate_2

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	var m = make(map[int][]int)
	for idx, n := range nums {
		if l, ok := m[n]; !ok {
			m[n] = []int{idx}
		} else {
			for _, i := range l {
				if (idx - i) <= k {
					return true
				}
			}
			l = append(l, idx)
			m[n] = l
		}
	}
	return false
}
