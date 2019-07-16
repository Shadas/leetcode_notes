package _217_Contains_Duplicate

func containsDuplicate(nums []int) bool {
	var m = make(map[int]bool)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		} else {
			m[n] = true
		}
	}
	return false
}
