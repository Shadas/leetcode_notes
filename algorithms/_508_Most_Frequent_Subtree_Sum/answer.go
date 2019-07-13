package _508_Most_Frequent_Subtree_Sum

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	var (
		m        = make(map[int]int)
		maxCount int
		ret      []int
	)
	push(m, root)
	for _, v := range m {
		if v > maxCount {
			maxCount = v
		}
	}
	for k, v := range m {
		if v == maxCount {
			ret = append(ret, k)
		}
	}
	return ret
}

func push(m map[int]int, n *TreeNode) (ret int) {
	if n == nil {
		return
	}
	var (
		left  int
		right int
	)
	if n.Left != nil {
		left = push(m, n.Left)
	}
	if n.Right != nil {
		right = push(m, n.Right)
	}
	ret = left + right + n.Val
	if c, ok := m[ret]; ok {
		m[ret] = c + 1
	} else {
		m[ret] = 1
	}
	return
}
