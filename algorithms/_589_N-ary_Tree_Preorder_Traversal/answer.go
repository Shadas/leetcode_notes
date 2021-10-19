package _589_N_ary_Tree_Preorder_Traversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	//return preOrderR(root)
	return preOrderI(root)
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

// 迭代解法
func preOrderI(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var (
		ret   []int
		stack = []*Node{root}
	)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		ret = append(ret, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return ret
}
