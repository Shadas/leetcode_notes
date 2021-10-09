package _297_Serialize_and_Deserialize_Binary_Tree

import (
	"fmt"
	"testing"
)

func TestPreOrderSerialize(t *testing.T) {
	var (
		root  *TreeNode
		str   string
		start int
	)
	root = &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2}}
	str = preOrderSerialize(root)
	fmt.Println(str)
	start = 0
	root = preOrderDeserialize(str, &start)
	fmt.Println(root)
	root = nil
	str = preOrderSerialize(root)
	fmt.Println(str)
	start = 0
	root = preOrderDeserialize(str, &start)
	fmt.Println(root)
}

func TestPostOrderSerialize(t *testing.T) {
	var (
		root  *TreeNode
		str   string
		start int
	)
	root = &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2}}
	str = postOrderSerialize(root)
	fmt.Println(str)
	start = len(str) - 1
	root = postOrderDeserialize(str, &start)
	fmt.Println(root)
	root = nil
	str = postOrderSerialize(root)
	fmt.Println(str)
	start = len(str) - 1
	root = postOrderDeserialize(str, &start)
	fmt.Println(root)
}

func TestBFSSerialize(t *testing.T) {
	var (
		root *TreeNode
		str  string
	)
	root = &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2}}
	str = BFSSerialize(root)
	fmt.Println(str)
	root = BFSDeserialize(str)
	fmt.Println(root, root.Left, root.Right)
	root = nil
	str = BFSSerialize(root)
	fmt.Println(str)
	root = BFSDeserialize(str)
	fmt.Println(root)
}
