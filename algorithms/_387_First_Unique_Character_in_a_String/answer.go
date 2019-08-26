package _387_First_Unique_Character_in_a_String

func firstUniqChar(s string) int {
	var fm = make(map[rune]int)
	for _, b := range s {
		count, _ := fm[b]
		count++
		fm[b] = count
	}
	for idx, b := range s {
		if fm[b] == 1 {
			return idx
		}
	}
	return -1
}
