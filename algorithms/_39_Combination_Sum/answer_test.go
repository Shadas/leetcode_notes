package _39_Combination_Sum

import "testing"

func TestCombinationSum(t *testing.T) {
	var (
		candidates []int
		target     int
		ret        [][]int
	)
	candidates = []int{2, 3, 5}
	target = 8
	ret = combinationSum(candidates, target)
	t.Log(ret)
}
