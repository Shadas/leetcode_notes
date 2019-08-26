package _242_Valid_Anagram

func isAnagram(s string, t string) bool {
	var ms, mt = make(map[rune]int), make(map[rune]int)
	for _, b := range s {
		count, _ := ms[b]
		count++
		ms[b] = count
	}
	for _, b := range t {
		count, _ := mt[b]
		count++
		mt[b] = count
	}
	for k, v := range ms {
		count, ok := mt[k]
		if !ok || count != v {
			return false
		}
	}
	for k, v := range mt {
		count, ok := ms[k]
		if !ok || count != v {
			return false
		}
	}
	return true
}
