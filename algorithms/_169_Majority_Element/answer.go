package _169_Majority_Element

func majorityElement(nums []int) int {
	return majorityElementByMapCount(nums)
}

func majorityElementByMapCount(nums []int) int {
	l := len(nums)
	m := make(map[int]int)
	for _, n := range nums {
		var tmpC int
		if c, ok := m[n]; ok {
			tmpC = c + 1
		} else {
			tmpC = 1
		}
		if tmpC > l/2 {
			return n
		}
		m[n] = tmpC
	}
	return 0
}
