package _513_Find_Bottom_Left_Tree_Value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	//return findBottomLeftValueDFS(root)
	return findBottomLeftValueBFS(root)
}

// BFS 解法
func findBottomLeftValueBFS(root *TreeNode) int {
	if root == nil {
		return -1
	}
	var (
		level, tmp []*TreeNode
		first      = true
		ret        int
	)
	level = append(level, root)
	for len(level) != 0 {
		first = true
		for _, node := range level {
			if first {
				ret = node.Val
				first = false
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		level = tmp
		tmp = []*TreeNode{}
	}
	return ret
}

// DFS 解法
func findBottomLeftValueDFS(root *TreeNode) int {
	ret, _ := DFS(root, 0, 0)
	return ret
}

func DFS(root *TreeNode, value, depth int) (nv, nd int) {
	if root == nil {
		return value, depth
	}
	nvl, ndl := DFS(root.Left, root.Val, depth+1)
	nvr, ndr := DFS(root.Right, root.Val, depth+1)
	if ndl >= ndr {
		nd = ndl
		nv = nvl
	} else {
		nd = ndr
		nv = nvr
	}
	return
}
