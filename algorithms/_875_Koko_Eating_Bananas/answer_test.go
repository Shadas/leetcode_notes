package _875_Koko_Eating_Bananas

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	if ret := minEatingSpeed([]int{3, 6, 7, 11}, 8); ret != 4 {
		t.Errorf("ret is %d, not 4", ret)
	}
	if ret := minEatingSpeed([]int{30, 11, 23, 4, 20}, 6); ret != 23 {
		t.Errorf("ret is %d, not 23", ret)
	}
	if ret := minEatingSpeed([]int{2, 2}, 6); ret != 1 {
		t.Errorf("ret is %d, not 1", ret)
	}
	if ret := minEatingSpeed([]int{312884470}, 312884469); ret != 2 {
		t.Errorf("ret is %d, not 2", ret)
	}
}

func TestEatOver(t *testing.T) {
	if ret := eatOver([]int{30, 11, 23, 4, 20}, 23, 6); ret == false {
		t.Errorf("ret need be true")
	}
}
