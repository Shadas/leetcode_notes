package _229_Majority_Element_2

func majorityElement(nums []int) []int {
	return majorityElementByCount(nums)
}

func majorityElementByCount(nums []int) []int {
	l := len(nums)
	m := make(map[int]int)
	ret := []int{}

	for _, n := range nums {
		var tmpC int
		if c, ok := m[n]; ok {
			tmpC = c + 1
		} else {
			tmpC = 1
		}
		m[n] = tmpC
	}

	for n, c := range m {
		if c > l/3 {
			ret = append(ret, n)
		}
	}
	return ret
}
