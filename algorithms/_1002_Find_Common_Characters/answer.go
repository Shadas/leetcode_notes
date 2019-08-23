package _1002_Find_Common_Characters

func commonChars(A []string) []string {
	if len(A) == 0 {
		return []string{}
	}
	var collections []map[string]int
	for _, str := range A {
		tmpCollection := make(map[string]int)
		for _, b := range str {
			if count, ok := tmpCollection[string(b)]; ok {
				count++
				tmpCollection[string(b)] = count
			} else {
				tmpCollection[string(b)] = 1
			}
		}
		collections = append(collections, tmpCollection)
	}
	var m = collections[0]
	for _, tm := range collections {
		m = mergeMinMap(m, tm)
	}
	var ret []string
	for k, v := range m {
		for i := 0; i < v; i++ {
			ret = append(ret, k)
		}
	}
	return ret
}

func mergeMinMap(m1, m2 map[string]int) (ret map[string]int) {
	ret = make(map[string]int)
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; ok {
			var v int
			if v1 > v2 {
				v = v2
			} else {
				v = v1
			}
			ret[k1] = v
		}
	}
	return
}
