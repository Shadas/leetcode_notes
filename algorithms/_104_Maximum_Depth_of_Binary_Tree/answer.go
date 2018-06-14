package _104_Maximum_Depth_of_Binary_Tree

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

func maxDepth(root *TreeNode) int {
	// return maxDepthRecursively(root)
	return maxDepthIteratively(root)
}

// recursively
func maxDepthRecursively(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepthRecursively(root.Left), maxDepthRecursively(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// iteratively
func maxDepthIteratively(root *TreeNode) (ret int) {
	if root == nil {
		return 0
	}
	queue := Queue{}
	queue.Push(root)
	for {
		if queue.IsEmpty() {
			break
		}
		ret += 1
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Pop()
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
	}
	return
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
