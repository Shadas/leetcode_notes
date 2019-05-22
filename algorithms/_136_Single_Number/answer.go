package _136_Single_Number

func singleNumber(nums []int) int {
	return singleNumber2(nums)
}

func singleNumber1(nums []int) int {
	var m = make(map[int]int)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			delete(m, n)
		} else {
			m[n] = 0
		}
	}
	for r, _ := range m {
		return r
	}
	return 0
}
