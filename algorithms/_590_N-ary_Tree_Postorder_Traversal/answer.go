package _590_N_ary_Tree_Postorder_Traversal

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	return postorderR(root)
}

// 递归解法
func postorderR(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var (
		ret []int
	)
	for _, node := range root.Children {
		ret = append(ret, postorderR(node)...)
	}
	ret = append(ret, root.Val)
	return ret
}
