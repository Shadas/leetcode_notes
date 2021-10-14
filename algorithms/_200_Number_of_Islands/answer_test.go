package _200_Number_of_Islands

import "testing"

func TestNumIslands(t *testing.T) {
	var (
		grid  [][]byte
		count int
	)

	grid = [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	if count = numIslands(grid); count != 1 {
		t.Errorf("wrong count with %d", count)
	}
}
