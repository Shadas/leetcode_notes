package _520_Detect_Capital

func detectCapitalUse(word string) bool {
	var (
		allCapital       = true
		allNotCapital    = true
		onlyFirstCapital = true
	)

	for idx, b := range word {
		if idx == 0 {
			if b >= 'a' && b <= 'z' {
				onlyFirstCapital = false
				allCapital = false
			} else {
				allNotCapital = false
			}
		} else {
			if allCapital && b >= 'a' && b <= 'z' {
				allCapital = false
			}
			if allNotCapital && b >= 'A' && b <= 'Z' {
				allNotCapital = false
			}
			if onlyFirstCapital && b >= 'A' && b <= 'Z' {
				onlyFirstCapital = false
			}
		}
		if !allCapital && !allNotCapital && !onlyFirstCapital {
			return false
		}
	}
	return allCapital || allNotCapital || onlyFirstCapital
}
