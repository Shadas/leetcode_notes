package array

import (
	"testing"
)

func TestIsIntArrayEqual(t *testing.T) {
	a1 := []int{1, 2, 3}
	a2 := []int{1, 2, 3}
	a3 := []int{1, 1, 3}

	if ret := IsIntArrayEqual(a1, a2); !ret {
		t.Error("is equal error.")
	}
	if ret := IsIntArrayEqual(a1, a3); ret {
		t.Error("is equal error.")
	}
}
