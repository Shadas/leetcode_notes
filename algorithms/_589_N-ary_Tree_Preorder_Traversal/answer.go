package _589_N_ary_Tree_Preorder_Traversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	return preOrderR(root)
}

// 递归解法
func preOrderR(root *Node) []int {
	var ret []int
	if root == nil {
		return ret
	}
	ret = append(ret, root.Val)
	for _, n := range root.Children {
		ret = append(ret, preOrderR(n)...)
	}
	return ret
}
