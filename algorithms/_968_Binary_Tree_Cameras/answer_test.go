package _968_Binary_Tree_Cameras

import (
	"fmt"
	"testing"
)

func TestMinCameraCover(t *testing.T) {
	root := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Left: &TreeNode{
					Left:  nil,
					Right: &TreeNode{},
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	}
	count := minCameraCover(root)
	fmt.Println(count)
}
