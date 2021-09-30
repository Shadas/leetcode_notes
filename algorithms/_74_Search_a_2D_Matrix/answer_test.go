package _74_Search_a_2D_Matrix

import "testing"

func TestIdxToPos(t *testing.T) {
	var (
		i, j int
	)
	if i, j = idxToPos(0, 10); i != 0 && j != 0 {
		t.Errorf("wrong ret with i=%d, j=%d", i, j)
	}
}

func TestSearchMatrix(t *testing.T) {
	var (
		matrix [][]int
		target int
	)
	matrix, target = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3
	if !searchMatrix(matrix, target) {
		t.Errorf("should be true")
	}
	matrix, target = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13
	if searchMatrix(matrix, target) {
		t.Errorf("should be false")
	}
	matrix, target = [][]int{{1}}, 1
	if !searchMatrix(matrix, target) {
		t.Errorf("should be true")
	}
	matrix, target = [][]int{{1, 3}}, 3
	if !searchMatrix(matrix, target) {
		t.Errorf("should be true")
	}
}
