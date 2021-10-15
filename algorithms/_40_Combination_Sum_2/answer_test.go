package _40_Combination_Sum_2

import "testing"

func TestCombinationSum(t *testing.T) {
	var (
		candidates []int
		target     int
		ret        [][]int
	)
	candidates = []int{10, 1, 2, 7, 6, 1, 5}
	target = 8
	ret = combinationSum2(candidates, target)
	t.Log(ret)
	candidates = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	target = 10
	ret = combinationSum2(candidates, target)
	t.Log(ret)
	//candidates = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	//target = 30
	//ret = combinationSum2(candidates, target)
	t.Log(ret)
}
