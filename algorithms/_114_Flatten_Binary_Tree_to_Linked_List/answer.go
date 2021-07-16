package _114_Flatten_Binary_Tree_to_Linked_List

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	//flattenRecursion(root)
	flattenIteration(root)
}

func flattenIteration(root *TreeNode) {
	cur := root
	for cur != nil {
		lp := cur.Left
		if cur.Left != nil {
			for lp.Right != nil {
				lp = lp.Right
			}
			lp.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
}

func flattenRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		flattenRecursion(root.Left)
	}
	if root.Right != nil {
		flattenRecursion(root.Right)
	}
	tmp := root.Right
	root.Right = root.Left
	root.Left = nil
	x := root
	for x.Right != nil {
		x = x.Right
	}
	x.Right = tmp
}
