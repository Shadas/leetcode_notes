package _429_N_ary_Tree_Level_Order_Traversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var (
		ret   [][]int
		q, tq []*Node
	)
	if root == nil {
		return ret
	}
	q = append(q, root)
	for len(q) > 0 {
		tr := []int{}
		for _, n := range q {
			if n == nil {
				continue
			}
			tr = append(tr, n.Val)
			tq = append(tq, n.Children...)
		}
		q = tq
		tq = []*Node{}
		ret = append(ret, tr)
	}
	return ret
}
