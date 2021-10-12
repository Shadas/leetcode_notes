package _1260_Shift_2D_Grid

import (
	"fmt"
	"testing"
)

func TestShiftGrid(t *testing.T) {
	var (
		grid [][]int
		k    int
		ret  [][]int
	)
	grid, k = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1
	ret = shiftGrid(grid, k)
	//fmt.Println(ret)

	grid, k = [][]int{{1}}, 100
	ret = shiftGrid(grid, k)
	fmt.Println(ret)
}
