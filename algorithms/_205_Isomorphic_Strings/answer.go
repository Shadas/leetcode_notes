package _205_Isomorphic_Strings

func isIsomorphic(s string, t string) bool {
	var pms, pmt = make(map[rune][]int), make(map[rune][]int)
	for idx, b := range s {
		list, _ := pms[b]
		list = append(list, idx)
		pms[b] = list
	}
	for _, idxs := range pms {
		b := t[idxs[0]]
		for _, idx := range idxs {
			if b != t[idx] {
				return false
			}
		}
	}
	for idx, b := range t {
		list, _ := pmt[b]
		list = append(list, idx)
		pmt[b] = list
	}
	for _, idxs := range pmt {
		b := s[idxs[0]]
		for _, idx := range idxs {
			if b != s[idx] {
				return false
			}
		}
	}
	return true
}
