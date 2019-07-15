package _108_Convert_Sorted_Array_to_Binary_Search_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTRecursion(nums)
}

func sortedArrayToBSTRecursion(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var (
		mid int
		ret *TreeNode
	)
	mid = len(nums) / 2
	ret = &TreeNode{Val: nums[mid]}
	ret.Left = sortedArrayToBSTRecursion(nums[:mid])
	if mid >= len(nums)-1 {
		ret.Right = nil
	} else {
		ret.Right = sortedArrayToBSTRecursion(nums[mid+1:])
	}
	return ret
}
