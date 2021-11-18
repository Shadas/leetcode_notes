package _559_Maximum_Depth_of_N_ary_Tree

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	return maxDepthRecursively(root)
}

// 递归解法
func maxDepthRecursively(root *Node) int {
	if root == nil {
		return 0
	}
	var (
		result int
	)
	for _, child := range root.Children {
		tr := maxDepthRecursively(child)
		if tr > result {
			result = tr
		}
	}
	return result + 1
}
