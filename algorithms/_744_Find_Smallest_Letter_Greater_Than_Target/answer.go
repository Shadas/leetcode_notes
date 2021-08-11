package _744_Find_Smallest_Letter_Greater_Than_Target

func nextGreatestLetter(letters []byte, target byte) byte {
	return findIteration(letters, target)
}

func findIteration(letters []byte, target byte) byte {
	left, right := 0, len(letters)
	for left < right {
		mid := left + (right-left)/2
		if letters[mid] < target+1 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if right == len(letters) {
		return letters[0]
	} else {
		return letters[right]
	}
}
