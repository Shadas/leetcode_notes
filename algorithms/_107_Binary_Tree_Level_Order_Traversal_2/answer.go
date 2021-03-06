package _107_Binary_Tree_Level_Order_Traversal_2

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var ret = [][]int{}
	var queue = Queue{}
	if root == nil {
		return ret
	}
	queue.Push(root)
	for {
		if queue.IsEmpty() {
			break
		}
		tmp := []int{}
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Pop()
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		ret = append(ret, tmp)
	}
	var ret1 = [][]int{}
	for i := len(ret); i > 0; i-- {
		ret1 = append(ret1, ret[i-1])
	}
	return ret1
}

type Queue struct {
	queue []*TreeNode
}

func (q *Queue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue) Push(node *TreeNode) {
	q.queue = append(q.queue, node)
}

func (q *Queue) Pop() (node *TreeNode) {
	if len(q.queue) == 0 {
		return
	}
	node = q.queue[0]
	q.queue = q.queue[1:]
	return
}

func (q *Queue) Len() int {
	return len(q.queue)
}
