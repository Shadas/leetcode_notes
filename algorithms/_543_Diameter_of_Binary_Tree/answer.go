package _543_Diameter_of_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	diameterOfBinaryTreeR(root)
	return max
}

func diameterOfBinaryTreeR(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := diameterOfBinaryTreeR(root.Left)
	r := diameterOfBinaryTreeR(root.Right)
	if l+r > max {
		max = l + r
	}
	curMax := l
	if r > l {
		curMax = r
	}
	return curMax + 1
}
