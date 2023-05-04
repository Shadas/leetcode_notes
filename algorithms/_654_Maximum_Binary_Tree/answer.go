package _654_Maximum_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	//return constructMaximumBinaryTreeRecursively(nums)
	return constructMaximumBinaryTreeStack(nums)
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

func constructMaximumBinaryTreeStack(nums []int) *TreeNode {
	var (
		treeNodePos = make([]*TreeNode, len(nums))
		stack       []int // stack of num index
		maxPosIdx   = 0
	)
	for idx, num := range nums {
		currNode := &TreeNode{
			Val: num,
		}
		treeNodePos[idx] = currNode
		for len(stack) != 0 {
			if nums[stack[len(stack)-1]] < num {
				currNode.Left = treeNodePos[stack[len(stack)-1]]
				// pop
				stack = stack[0 : len(stack)-1]
			} else {
				topIdx := stack[len(stack)-1]
				treeNodePos[topIdx].Right = currNode
				stack = append(stack, idx)
				break
			}
		}
		if len(stack) == 0 {
			stack = append(stack, idx)
			maxPosIdx = idx
		}
	}
	return treeNodePos[maxPosIdx]
}
