package _350_Intersection_of_Two_Arrays_2

func intersect(nums1 []int, nums2 []int) []int {
	var m1, m2 = make(map[int]int), make(map[int]int)
	for _, n := range nums1 {
		if count, ok := m1[n]; ok {
			count++
			m1[n] = count
		} else {
			m1[n] = 1
		}
	}
	for _, n := range nums2 {
		if count, ok := m2[n]; ok {
			count++
			m2[n] = count
		} else {
			m2[n] = 1
		}
	}
	var ret = make(map[int]int)
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok {
			if v1 > v2 {
				ret[k1] = v2
			} else {
				ret[k1] = v1
			}
		}
	}
	var r []int
	for k, v := range ret {
		for i := 0; i < v; i++ {
			r = append(r, k)
		}
	}
	return r
}
