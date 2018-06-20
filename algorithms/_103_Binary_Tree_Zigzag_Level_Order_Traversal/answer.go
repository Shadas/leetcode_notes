package _103_Binary_Tree_Zigzag_Level_Order_Traversal

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

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ret = [][]int{}
	var queue = Queue{}
	var direction bool
	if root == nil {
		return ret
	}
	queue.Push(root)
	for {
		direction = !direction
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
		if direction {
			ret = append(ret, tmp)
		} else {
			tmp1 := []int{}
			for i := len(tmp); i > 0; i-- {
				tmp1 = append(tmp1, tmp[i-1])
			}
			ret = append(ret, tmp1)
		}
	}
	return ret
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
