package _51_N_Queens

import "testing"

func TestSolveNQueens(t *testing.T) {
	var ret [][]string
	ret = solveNQueens(1)
	t.Log(ret)
	ret = solveNQueens(4)
	t.Log(ret)

}
