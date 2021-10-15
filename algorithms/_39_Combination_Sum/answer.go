package _39_Combination_Sum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var (
		result [][]int
	)
	sort.Ints(candidates)
	combinationSumR(candidates, target, 0, []int{}, &result)
	return result
}

func combinationSumR(candidates []int, target, idx int, path []int, result *[][]int) {
	if target == 0 {
		x := make([]int, len(path))
		for idx, v := range path {
			x[idx] = v
		}
		*result = append(*result, x)
		return
	}
	if len(path) != 0 && path[len(path)-1] > target {
		return
	}
	for i := idx; i < len(candidates); i++ {
		x := candidates[i]
		combinationSumR(candidates, target-x, i, append(path, x), result)
	}
}
