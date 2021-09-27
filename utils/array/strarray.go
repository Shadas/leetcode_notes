package array

// IsStrArrayEqual compares two []string and returns if they are the same.
func IsStrArrayEqual(a1, a2 []string) bool {
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
