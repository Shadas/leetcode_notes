package _349_Intersection_of_Two_Arrays

func intersection(nums1 []int, nums2 []int) []int {
	var m1, m2 = make(map[int]interface{}), make(map[int]interface{})
	for _, n := range nums1 {
		m1[n] = nil
	}
	for _, n := range nums2 {
		m2[n] = nil
	}
	var ret []int
	for n1, _ := range m1 {
		if _, ok := m2[n1]; ok {
			ret = append(ret, n1)
		}
	}
	return ret
}
