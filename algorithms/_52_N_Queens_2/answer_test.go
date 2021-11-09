package _52_N_Queens_2

import "testing"

func TestTotalNQueens(t *testing.T) {
	if ret := totalNQueens(1); ret != 1 {
		t.Errorf("wrong ret with %d", ret)
	}
	if ret := totalNQueens(4); ret != 2 {
		t.Errorf("wrong ret with %d", ret)
	}
}
