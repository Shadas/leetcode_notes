package _106_Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

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

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderLen := len(inorder)
	postLen := len(postorder)
	return buildTreeRecursion(inorder, 0, inorderLen-1, postorder, 0, postLen-1)
}

func buildTreeRecursion(inorder []int, inStart int, inEnd int, postorder []int, postStart int, postEnd int) *TreeNode {
	if inStart > inEnd || postStart > postEnd {
		return nil
	}

	var rootVal = postorder[postEnd]
	var rootIndex = 0
	for i := 0; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}
	length := rootIndex - inStart // 左子树长度
	node := &TreeNode{
		Val: rootVal,
	}
	node.Right = buildTreeRecursion(inorder, rootIndex+1, inEnd, postorder, postStart+length, postEnd-1)
	node.Left = buildTreeRecursion(inorder, inStart, rootIndex-1, postorder, postStart, postStart+length-1)

	return node
}
