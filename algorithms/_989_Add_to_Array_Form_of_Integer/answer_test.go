package _989_Add_to_Array_Form_of_Integer

import "testing"

func isSame(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestAddToArrayForm(t *testing.T) {
	if ret := addToArrayForm([]int{0}, 23); !isSame(ret, []int{2, 3}) {
		t.Errorf("wrong with %v", ret)
	}
}
