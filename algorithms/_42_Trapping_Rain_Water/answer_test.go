package _42_Trapping_Rain_Water

import "testing"

func TestTrap(t *testing.T) {
	if sum := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}); sum != 6 {
		t.Errorf("wrong number is %d", sum)
	}
}
