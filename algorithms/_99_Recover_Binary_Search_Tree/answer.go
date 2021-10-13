package _99_Recover_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	pre    *TreeNode
	m1, m2 *TreeNode
)

func recoverTree(root *TreeNode) {
	pre, m1, m2 = nil, nil, nil
	dfs(root)
	m1.Val, m2.Val = m2.Val, m1.Val
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if pre != nil && pre.Val > root.Val {
		if m1 == nil {
			m1 = pre
		}
		m2 = root
	}
	pre = root
	dfs(root.Right)
	return
}
