package _257_Binary_Tree_Paths

import (
	"fmt"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	var ret []string
	binaryTreePathsRecursion(root, []string{}, &ret)
	return ret
}

func binaryTreePathsRecursion(node *TreeNode, trace []string, ret *[]string) {
	trace = append(trace, fmt.Sprint(node.Val))
	if node.Left == nil && node.Right == nil {
		*ret = append(*ret, strings.Join(trace, "->"))
		return
	}
	if node.Left != nil {
		binaryTreePathsRecursion(node.Left, trace, ret)
	}
	if node.Right != nil {
		binaryTreePathsRecursion(node.Right, trace, ret)
	}
}
