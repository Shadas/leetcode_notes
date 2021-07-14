package _968_Binary_Tree_Cameras

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type status int

const (
	installed status = 0
	seen      status = 1
	ignored   status = 2
)

var count int

func minCameraCover(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if x := getStatus(root); x == ignored {
		count++
	}
	return count
}

func getStatus(node *TreeNode) status {
	if node == nil {
		return seen
	}
	ls, rs := getStatus(node.Left), getStatus(node.Right)
	if ls == ignored || rs == ignored { // 子节点有任一不被监控，当前节点要安装
		count++
		return installed
	} else if ls == installed || rs == installed { // 子节点有任一安装了监控，则此节点必被监控
		return seen
	} else { // 俩都是seen，则此节点必不被监控
		return ignored
	}
}
