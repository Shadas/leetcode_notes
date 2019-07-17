package _383_Ransom_Note

func canConstruct(ransomNote string, magazine string) bool {
	var m map[rune]int = make(map[rune]int)
	for _, b := range magazine {
		if c, ok := m[b]; ok {
			m[b] = c + 1
		} else {
			m[b] = 1
		}
	}
	for _, b := range ransomNote {
		if c, ok := m[b]; ok {
			if c-1 < 0 {
				return false
			} else {
				m[b] = c - 1
			}
		} else {
			return false
		}
	}
	return true
}
