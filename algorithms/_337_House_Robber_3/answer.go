package _337_House_Robber_3

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	// return solution1(root)
	return solution2(root)
}

// 递归解法
func solution1(root *TreeNode) int {
	with, without := solution1Rec(root)
	return int(math.Max(float64(with), float64(without)))
}

// 返回 包含 或 不包含 该节点的值
func solution1Rec(root *TreeNode) (with, without int) {
	if root == nil {
		return 0, 0
	}
	lwith, lwithout := solution1Rec(root.Left)
	rwith, rwithout := solution1Rec(root.Right)
	with = root.Val + lwithout + rwithout
	without = int(math.Max(float64(lwith), float64(lwithout)) + math.Max(float64(rwith), float64(rwithout)))
	return
}

// 递归解法，带计算结果缓存
func solution2(root *TreeNode) int {
	cache := make(map[*TreeNode][2]int)
	ret := solution2Rec(root, cache)
	return int(math.Max(float64(ret[0]), float64(ret[1])))
}

func solution2Rec(root *TreeNode, cache map[*TreeNode][2]int) (ret [2]int) {
	if root == nil {
		return
	}
	var ok bool
	if ret, ok = cache[root]; ok {
		return
	}
	lret := solution2Rec(root.Left, cache)
	rret := solution2Rec(root.Right, cache)
	cache[root.Left] = lret
	cache[root.Right] = rret
	ret[0] = root.Val + lret[1] + rret[1]
	ret[1] = int(math.Max(float64(lret[0]), float64(lret[1])) + math.Max(float64(rret[0]), float64(rret[1])))
	cache[root] = ret
	return
}
