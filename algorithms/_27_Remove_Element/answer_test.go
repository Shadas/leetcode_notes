package _27_Remove_Element

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	if ret := removeElement([]int{3, 2, 2, 3}, 3); ret != 2 {
		t.Error("not 2 in test1.")
	}
}
