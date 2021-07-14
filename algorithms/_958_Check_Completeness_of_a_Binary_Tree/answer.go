package _958_Check_Completeness_of_a_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue = []*TreeNode{root}
	var meetNil bool
	for len(queue) != 0 {
		n := queue[0] // head
		queue = queue[1:]
		if n.Left == nil {
			meetNil = true
		} else {
			if meetNil {
				return false
			} else {
				queue = append(queue, n.Left)
			}
		}
		if n.Right == nil {
			meetNil = true
		} else {
			if meetNil {
				return false
			} else {
				queue = append(queue, n.Right)
			}
		}
	}
	return true
}
