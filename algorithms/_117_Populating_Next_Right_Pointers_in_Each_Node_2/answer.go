package _117_Populating_Next_Right_Pointers_in_Each_Node_2

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	return connectIteration(root)
}

// 层序迭代
func connectIteration(root *Node) *Node {
	if root == nil {
		return root
	}
	var queue = []*Node{root, nil}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur != nil {
			cur.Next = queue[0]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		} else {
			if len(queue) == 0 || queue[0] == nil {
				return root
			} else {
				queue = append(queue, nil)
			}
		}
	}
	return root
}
