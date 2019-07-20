package _796_Rotate_String

func rotateString(A string, B string) bool {
	if A == "" && B == "" {
		return true
	}
	if len(A) < len(B) {
		return false
	}
	for i, s := range B {
		if s == rune(A[0]) {
			tmpStr := B[i:] + B[0:i]
			if tmpStr == A {
				return true
			}
		}
	}
	return false
}
