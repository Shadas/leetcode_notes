package _344_Reverse_String

func reverseString(s []byte) {
	for idx := 0; idx < len(s); idx++ {
		s[idx], s[len(s)-1-idx] = s[len(s)-1-idx], s[idx]
		if idx >= len(s)/2-1 {
			break
		}
	}
}
