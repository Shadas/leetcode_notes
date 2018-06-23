package _105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	preLength := len(preorder)
	inLength := len(inorder)
	return buildTreeRecursion(preorder, 0, preLength-1, inorder, 0, inLength-1)
}

func buildTreeRecursion(preorder []int, preStart int, preEnd int, inorder []int, inStart int, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	var rootVal = preorder[preStart]
	var rootIndex = 0
	for i := inStart; i <= inEnd; i++ {
		if rootVal == inorder[i] {
			rootIndex = i
			break
		}
	}
	length := rootIndex - inStart
	node := &TreeNode{
		Val: rootVal,
	}
	node.Left = buildTreeRecursion(preorder, preStart+1, preStart+length, inorder, inStart, rootIndex-1)
	node.Right = buildTreeRecursion(preorder, preStart+length+1, preEnd, inorder, rootIndex+1, inEnd)

	return node
}
