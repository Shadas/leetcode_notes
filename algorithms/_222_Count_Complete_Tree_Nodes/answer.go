package _222_Count_Complete_Tree_Nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	return countNodesRecursion(root)
}

func countNodesRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodesRecursion(root.Left) + countNodesRecursion(root.Right)
}
