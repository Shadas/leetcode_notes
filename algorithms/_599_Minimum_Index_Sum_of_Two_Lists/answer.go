package _599_Minimum_Index_Sum_of_Two_Lists

func findRestaurant(list1 []string, list2 []string) []string {
	return findRestaurantLoopSearch(list1, list2)
}

func findRestaurantLoopSearch(list1 []string, list2 []string) []string {
	var (
		m          = make(map[int][]string)
		minIdx int = len(list1) + len(list2)
	)
	for k1, v1 := range list1 {
		for k2, v2 := range list2 {
			if v1 == v2 {
				sum := k1 + k2
				if l, ok := m[sum]; ok {
					l = append(l, v2)
					m[sum] = l
				} else {
					m[sum] = []string{v2}
				}
				break
			}
		}
	}
	for idx, _ := range m {
		if idx < minIdx {
			minIdx = idx
		}
	}
	return m[minIdx]
}
