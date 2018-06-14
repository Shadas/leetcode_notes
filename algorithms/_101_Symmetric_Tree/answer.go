package _101_Symmetric_Tree

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

func isSymmetric(root *TreeNode) bool {
	return isSymmetricRecursively(root)
	// return isSymmetricIteratively(root)
}

// recursively
func isSymmetricRecursively(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compareRoot(root.Left, root.Right)
}

func compareRoot(leftTree, rightTree *TreeNode) bool {
	if leftTree == nil {
		return rightTree == nil
	}
	if rightTree == nil {
		return false
	}
	if leftTree.Val != rightTree.Val {
		return false
	}

	return compareRoot(leftTree.Left, rightTree.Right) && compareRoot(leftTree.Right, rightTree.Left)
}

// iteratively
func isSymmetricIteratively(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if (root.Left == nil && root.Right != nil) || (root.Right == nil && root.Left != nil) {
		return false
	}
	leftStack := Stack{}
	rightStack := Stack{}
	leftStack.Push(root.Left)
	rightStack.Push(root.Right)
	for {
		if leftStack.IsEmpty() && rightStack.IsEmpty() { // quit the loop
			break
		}
		leftTree, rightTree := leftStack.Pop(), rightStack.Pop()
		if leftTree.Val != rightTree.Val {
			return false
		}
		if (leftTree.Left == nil && rightTree.Right != nil) ||
			(leftTree.Left != nil && rightTree.Right == nil) ||
			(leftTree.Right == nil && rightTree.Left != nil) ||
			(leftTree.Right != nil && rightTree.Left == nil) {
			return false
		}
		if leftTree.Right != nil {
			leftStack.Push(leftTree.Right)
			rightStack.Push(rightTree.Left)
		}

		if leftTree.Left != nil {
			leftStack.Push(leftTree.Left)
			rightStack.Push(rightTree.Right)
		}

	}
	return true
}

type Stack struct {
	stack []*TreeNode
}

func (s *Stack) Push(node *TreeNode) {
	s.stack = append(s.stack, node)
}

func (s *Stack) Pop() (node *TreeNode) {
	node = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Len() int {
	return len(s.stack)
}
