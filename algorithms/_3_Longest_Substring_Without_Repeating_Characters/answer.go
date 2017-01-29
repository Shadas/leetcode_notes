package _3_Longest_Substring_Without_Repeating_Characters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	bstr := []byte(s)
	max := 0
	for i, _ := range bstr {
		ventor := []byte{bstr[i]}
		if max >= len(bstr)-i {
			break
		}
		for j := i + 1; j <= len(bstr)-1; j++ {
			if IndexByte(ventor, bstr[j]) != -1 {
				length := len(ventor)
				if length >= max {
					max = length
				}
				break
			} else {
				ventor = append(ventor, bstr[j])
				length := len(ventor)
				if length >= max {
					max = length
				}
			}
		}
	}
	return max
}

func IndexByte(arr []byte, i byte) int {
	ret := -1
	for m, ar := range arr {
		if ar == i {
			ret = m
			break
		}
	}
	return ret
}
