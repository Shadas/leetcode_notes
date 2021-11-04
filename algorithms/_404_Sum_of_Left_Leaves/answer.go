package _404_Sum_of_Left_Leaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	// return sumOfLeftLeavesBFS(root)
	// return sumOfLeftLeavesDFSRecursion(root)
	return sumOfLeftLeavesDFSUnrecursion(root)
}

func sumOfLeftLeavesDFSUnrecursion(node *TreeNode) int {
	var (
		s   = []*TreeNode{node}
		ret int
	)
	if node == nil {
		return 0
	}
	for len(s) > 0 {
		x := s[len(s)-1]
		s = s[0 : len(s)-1]
		if x.Right != nil {
			s = append(s, x.Right)
		}
		if x.Left != nil {
			s = append(s, x.Left)
			if x.Left.Left == nil && x.Left.Right == nil {
				ret += x.Left.Val
			}
		}
	}
	return ret
}

func sumOfLeftLeavesDFSRecursion(node *TreeNode) int {
	var ret int
	if node == nil {
		return 0
	}
	if node.Left != nil {
		ret += sumOfLeftLeavesDFSRecursion(node.Left)
		if node.Left.Left == nil && node.Left.Right == nil {
			ret += node.Left.Val
		}
	}
	if node.Right != nil {
		ret += sumOfLeftLeavesDFSRecursion(node.Right)
	}
	return ret
}

func sumOfLeftLeavesBFS(root *TreeNode) int {
	var (
		q   = []*TreeNode{root}
		ret int
	)
	if root == nil {
		return 0
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		if x.Left != nil {
			q = append(q, x.Left)
			if x.Left.Left == nil && x.Left.Right == nil {
				ret += x.Left.Val
			}
		}
		if x.Right != nil {
			q = append(q, x.Right)
		}
	}
	return ret
}
