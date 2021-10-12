package _134_Gas_Station

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	if ret := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}); ret != 3 {
		t.Errorf("wrong ret with %d", ret)
	}
}
