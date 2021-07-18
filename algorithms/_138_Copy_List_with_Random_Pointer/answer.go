package _138_Copy_List_with_Random_Pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	return copyRandomListDFS(head, nodeMap)
}

func copyRandomListDFS(node *Node, nodeMap map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, ok := nodeMap[node]; !ok {
		tmp := &Node{Val: node.Val}
		nodeMap[node] = tmp
		tmp.Next = copyRandomListDFS(node.Next, nodeMap)
		tmp.Random = copyRandomListDFS(node.Random, nodeMap)
	}
	return nodeMap[node]
}
