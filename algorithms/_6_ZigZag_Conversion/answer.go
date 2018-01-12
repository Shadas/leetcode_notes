package _6_ZigZag_Conversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var y = 0
	var flag = true
	var sarray = []string{}
	for i := 0; i < numRows; i++ {
		sarray = append(sarray, "")
	}
	for i := 0; i < len(s); i++ {
		sarray[y] += string(s[i])
		if y == 0 {
			flag = true
		}
		if y == numRows-1 {
			flag = false
		}
		if flag {
			y++
		} else {
			y--
		}
	}
	var ret string
	for i := 0; i < numRows; i++ {
		ret += sarray[i]
	}
	return ret
}
