package _47_Permutations_2

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	used := []bool{}
	for i := 0; i < len(nums); i++ {
		used = append(used, false)
	}
	backtrack(&ret, &[]int{}, nums, &used)
	return ret
}

func backtrack(list *[][]int, tmpList *[]int, nums []int, used *[]bool) {
	if len(*tmpList) == len(nums) {
		l := []int{}
		for _, v := range *tmpList {
			l = append(l, v)
		}
		*list = append(*list, l)
	} else {
		for i, n := range nums {
			if (*used)[i] || i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
				continue
			}
			(*used)[i] = true
			*tmpList = append(*tmpList, n)
			backtrack(list, tmpList, nums, used)
			(*used)[i] = false
			*tmpList = (*tmpList)[:len(*tmpList)-1]
		}
	}
}
