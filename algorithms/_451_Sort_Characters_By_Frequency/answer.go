package _451_Sort_Characters_By_Frequency

func frequencySort(s string) string {
	var (
		freq = make(map[rune]int)
		max  int
		ret  string
	)
	for _, r := range s {
		count, ok := freq[r]
		if !ok {
			count = 1
		} else {
			count++
		}
		if count >= max {
			max = count
		}
		freq[r] = count
	}
	for i := max; i >= 0; i-- {
		for b, v := range freq {
			if v == i {
				for j := 0; j < i; j++ {
					ret += string(b)
				}
			}
		}
	}
	return ret
}
