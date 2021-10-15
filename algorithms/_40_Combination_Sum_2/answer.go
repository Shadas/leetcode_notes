package _40_Combination_Sum_2

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var (
		result  [][]int
		resultM = make(map[string]struct{})
	)
	sort.Ints(candidates)
	combinationSumR(candidates, target, 0, []int{}, &result, resultM)
	return result
}

func combinationSumR(candidates []int, target int, idx int, path []int, result *[][]int, resultM map[string]struct{}) {
	if target < 0 {
		return
	}
	if target == 0 {
		x := make([]int, len(path))
		key := ""
		for idx, v := range path {
			x[idx] = v
			key += fmt.Sprintf("%d,", v)
		}
		if _, ok := resultM[key]; !ok {
			*result = append(*result, x)
			resultM[key] = struct{}{}
		}
		return
	}
	for i := idx; i < len(candidates); i++ {
		x := candidates[i]
		if i > idx && candidates[i] == candidates[i-1] {
			continue
		}
		combinationSumR(candidates, target-x, i+1, append(path, x), result, resultM)
	}
}
