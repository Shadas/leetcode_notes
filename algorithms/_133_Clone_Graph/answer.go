package _133_Clone_Graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	dict := make(map[int]*Node)
	return cloneGraphDFS(node, dict)
}

func cloneGraphDFS(node *Node, dict map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	tmp, ok := dict[node.Val]
	if ok {
		return tmp
	}
	tmp = &Node{
		Val: node.Val,
	}
	dict[node.Val] = tmp
	for _, n := range node.Neighbors {
		tmp.Neighbors = append(tmp.Neighbors, cloneGraphDFS(n, dict))
	}
	return tmp
}
