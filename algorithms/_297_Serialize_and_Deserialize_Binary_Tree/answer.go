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
	return postOrderSerialize(root)
	//return preOrderSerialize(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	//startCur := 0
	//return preOrderDeserialize(data, &startCur)
	startCur := len(data) - 1
	return postOrderDeserialize(data, &startCur)
}

// 后续遍历序列化
func postOrderSerialize(root *TreeNode) string {
	var str string
	if root == nil {
		str += ","
		return str
	}
	str += postOrderSerialize(root.Left)
	str += postOrderSerialize(root.Right)
	str += fmt.Sprintf(",%d", root.Val)
	return str
}

// 后序遍历反序列化
func postOrderDeserialize(data string, cur *int) *TreeNode {
	if string(data[*cur]) == "," {
		*cur -= 1
		return nil
	}
	begin := *cur
	for begin >= 0 && string(data[begin]) != "," {
		begin--
	}
	num, _ := strconv.ParseInt(data[begin+1:*cur+1], 10, 64)
	node := &TreeNode{Val: int(num)}
	*cur = begin - 1
	node.Right = postOrderDeserialize(data, cur)
	node.Left = postOrderDeserialize(data, cur)
	return node
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

// 先序遍历反序列化
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

//// 层序遍历序列化
//func bfsSerialize(root *TreeNode) string {
//	var (
//		str   string
//		level = []*TreeNode{root}
//	)
//	for len(level) != 0 {
//		node := level[0]
//		newStr := ""
//		if node == nil {
//			newStr = fmt.Sprintf("%d,", node.Val)
//		} else {
//			newStr = fmt.Sprintf("#,")
//		}
//		str += newStr
//	}
//	return str
//}
//
//// 层序遍历反序列化
//func bfsDeserialize(root *TreeNode) string {
//
//}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
