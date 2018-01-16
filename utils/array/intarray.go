package array

// IsIntArrayEqual compares two []int and returns if they are the same.
func IsIntArrayEqual(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
