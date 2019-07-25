package _501_Find_Mode_in_Binary_Search_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	var (
		ret     = []int{}
		maxFreq int
		m       = make(map[int]int)
		s       = []*TreeNode{}
		node    = root
	)
	for node != nil || len(s) > 0 {
		for node != nil {
			s = append(s, node)
			node = node.Left
		}
		if len(s) > 0 {
			node = s[len(s)-1]
			s = s[:len(s)-1]
			tmp := node.Val
			var (
				c  int
				ok bool
			)
			if c, ok = m[tmp]; ok {
				c = c + 1
			} else {
				c = 1
			}
			m[tmp] = c
			if c >= maxFreq {
				maxFreq = c
			}
			node = node.Right
		}
	}
	for k, v := range m {
		if v == maxFreq {
			ret = append(ret, k)
		}
	}
	return ret
}
