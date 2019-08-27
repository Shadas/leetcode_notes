package _49_Group_Anagrams

type Format struct {
	Fm    map[rune]int
	Count int
}

func groupAnagrams(strs []string) [][]string {
	var (
		formats  = make(map[string]Format)
		ret      [][]string
		emptyRet []string
	)
	for _, str := range strs {
		if str == "" {
			emptyRet = append(emptyRet, "")
			continue
		}
		tm, ok := formats[str]
		if !ok {
			tm = Format{
				Fm: make(map[rune]int),
			}
			for _, b := range str {
				count, _ := tm.Fm[b]
				count++
				tm.Fm[b] = count
			}
		}
		tm.Count++
		formats[str] = tm
	}
loop:
	for str, format := range formats {
		for idx, list := range ret {
			f0 := formats[list[0]]
			if isSame(f0.Fm, format.Fm) {
				for i := 0; i < format.Count; i++ {
					list = append(list, str)
				}
				ret[idx] = list
				continue loop
			}
		}
		newline := []string{}
		for i := 0; i < format.Count; i++ {
			newline = append(newline, str)
		}
		ret = append(ret, newline)
	}
	if len(emptyRet) != 0 {
		ret = append(ret, emptyRet)
	}
	return ret
}

func isSame(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		value, ok := b[k]
		if !ok || v != value {
			return false
		}
	}
	return true
}
