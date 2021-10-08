package _513_Find_Bottom_Left_Tree_Value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	return findBottomLeftValueDFS(root)
}

func findBottomLeftValueDFS(root *TreeNode) int {
	ret, _ := DFS(root, 0, 0)
	return ret
}

func DFS(root *TreeNode, value, depth int) (nv, nd int) {
	if root == nil {
		return value, depth
	}
	nvl, ndl := DFS(root.Left, root.Val, depth+1)
	nvr, ndr := DFS(root.Right, root.Val, depth+1)
	if ndl >= ndr {
		nd = ndl
		nv = nvl
	} else {
		nd = ndr
		nv = nvr
	}
	return
}
