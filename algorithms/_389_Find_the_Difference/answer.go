package _389_Find_the_Difference

func findTheDifference(s string, t string) byte {
	return findTheDifferenceWithXOR(s, t)
}

func findTheDifferenceWithXOR(s string, t string) byte {
	var (
		str = s + t
		ret rune
	)
	for _, b := range str {
		ret ^= b
	}
	return byte(ret)
}
