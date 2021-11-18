package _491_Increasing_Subsequences

import "fmt"

func findSubsequences(nums []int) [][]int {
	return findSubsequencesDFS(nums)
}

func findSubsequencesDFS(nums []int) [][]int {
	var (
		ss  [][]int
		nss [][]int
		sm  = make(map[string][]int)
	)
	dfs(nums, []int{}, &ss)
	for _, s := range ss {
		key := fmt.Sprint(s)
		if _, ok := sm[key]; !ok {
			sm[key] = s
			nss = append(nss, s)
		}
	}

	return nss
}

func dfs(nums []int, tmp []int, ss *[][]int) {
	if len(nums) == 0 {
		return
	}
	for idx, num := range nums {
		var newNums []int
		if idx != len(nums)-1 {
			newNums = nums[idx+1:]
		}
		if len(tmp) == 0 {
			dfs(newNums, append(tmp, num), ss)
		} else {
			if num >= tmp[len(tmp)-1] {
				x := make([]int, len(tmp))
				copy(x, tmp)
				x = append(x, num)
				*ss = append(*ss, x)
				dfs(newNums, append(tmp, num), ss)
			}
		}
	}
}
