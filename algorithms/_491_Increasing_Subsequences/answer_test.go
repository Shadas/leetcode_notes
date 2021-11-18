package _491_Increasing_Subsequences

import "testing"

func TestFindSubsequences(t *testing.T) {
	var (
		subsequences [][]int
	)
	subsequences = findSubsequences([]int{4, 6, 7, 7})
	t.Log(subsequences) // [[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
	subsequences = findSubsequences([]int{4, 4, 3, 2, 1})
	t.Log(subsequences) // [[4,4]]
	subsequences = findSubsequences([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	t.Log(subsequences) //
}
