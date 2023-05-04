package _449_Serialize_and_Deserialize_BST

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var (
		nodeValue []string
		line      []*TreeNode
	)
	if root != nil {
		line = append(line, root)
		nodeValue = append(nodeValue, strconv.FormatInt(int64(root.Val), 10))
	} else {
		nodeValue = append(nodeValue, "nil")
	}
	for len(line) != 0 {
		var tmpLine []*TreeNode
		for _, node := range line {
			if node.Left == nil {
				nodeValue = append(nodeValue, "nil")
			} else {
				nodeValue = append(nodeValue, strconv.FormatInt(int64(node.Left.Val), 10))
				tmpLine = append(tmpLine, node.Left)
			}
			if node.Right == nil {
				nodeValue = append(nodeValue, "nil")
			} else {
				nodeValue = append(nodeValue, strconv.FormatInt(int64(node.Right.Val), 10))
				tmpLine = append(tmpLine, node.Right)
			}
		}
		line = tmpLine
	}
	return strings.Join(nodeValue, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodeValues := strings.Split(data, ",")
	if len(nodeValues) == 1 && nodeValues[0] == "nil" {
		return nil
	}
	rootVal, _ := strconv.Atoi(nodeValues[0])
	root := &TreeNode{Val: rootVal}
	line := []*TreeNode{root}
	curIdx := 1
Finish:
	for len(line) != 0 {
		var tmpLine []*TreeNode
		for _, node := range line {
			if curIdx >= len(nodeValues) {
				break Finish
			}
			leftx := nodeValues[curIdx]
			if leftx == "nil" {
				node.Left = nil
			} else {
				leftValue, _ := strconv.Atoi(leftx)
				node.Left = &TreeNode{Val: leftValue}
				tmpLine = append(tmpLine, node.Left)
			}
			curIdx++
			if curIdx >= len(nodeValues) {
				break Finish
			}
			rightx := nodeValues[curIdx]
			if rightx == "nil" {
				node.Right = nil
			} else {
				rightValue, _ := strconv.Atoi(rightx)
				node.Right = &TreeNode{Val: rightValue}
				tmpLine = append(tmpLine, node.Right)
			}
			curIdx++
		}
		line = tmpLine
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
