package _95_Unique_Binary_Search_Trees_2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return genRecursive(1, n)
}

func genRecursive(i, j int) []*TreeNode {
	var ret = []*TreeNode{}
	if i > j {
		ret = append(ret, nil)
		return ret
	}
	if i == j {
		ret = append(ret, &TreeNode{Val: i})
		return ret
	}

	for k := i; k <= j; k++ {
		left := genRecursive(i, k-1)
		right := genRecursive(k+1, j)
		for _, tmpR := range right {
			for _, tmpL := range left {
				ret = append(ret, &TreeNode{
					Left:  tmpL,
					Right: tmpR,
					Val:   k,
				})
			}
		}
	}
	return ret
}
