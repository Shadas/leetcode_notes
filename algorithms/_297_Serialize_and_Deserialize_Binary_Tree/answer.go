package _297_Serialize_and_Deserialize_Binary_Tree

import (
	"fmt"
	"strconv"
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
	return preOrderSerialize(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	startCur := 0
	return preOrderDeserialize(data, &startCur)
}

// 先序遍历序列化
func preOrderSerialize(root *TreeNode) string {
	var str string
	if root == nil {
		str += ","
		return str
	}
	str += fmt.Sprintf("%d,", root.Val)
	str += preOrderSerialize(root.Left)
	str += preOrderSerialize(root.Right)
	return str
}

func preOrderDeserialize(data string, cur *int) *TreeNode {
	if string(data[*cur]) == "," { // 说明是空节点
		*cur += 1
		return nil
	}
	tail := *cur
	for tail < len(data) && string(data[tail]) != "," {
		tail++
	}
	num, _ := strconv.ParseInt(data[*cur:tail], 10, 64)
	node := &TreeNode{Val: int(num)}
	*cur = tail + 1
	node.Left = preOrderDeserialize(data, cur)
	node.Right = preOrderDeserialize(data, cur)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
