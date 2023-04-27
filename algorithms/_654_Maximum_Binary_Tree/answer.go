package _654_Maximum_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaximumBinaryTreeRecursively(nums)
}

func constructMaximumBinaryTreeRecursively(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// find maximum and split
	var (
		maxIdx int = 0
	)
	for idx, num := range nums {
		if num > nums[maxIdx] {
			maxIdx = idx
		}
	}
	x := &TreeNode{Val: nums[maxIdx]}
	x.Left = constructMaximumBinaryTreeRecursively(nums[0:maxIdx])
	x.Right = constructMaximumBinaryTreeRecursively(nums[maxIdx+1:])
	return x
}
