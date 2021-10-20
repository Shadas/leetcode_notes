package _208_Implement_Trie_Prefix_Tree

type Node struct {
	isStr bool      // 是否为完整字符串
	next  [26]*Node // a-z 位置可能的后继节点，如果为nil则尚未存在, idx=字符-'a'
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{root: &Node{}}
}

func (this *Trie) Insert(word string) {
	var cur *Node = this.root
	for _, b := range word {
		idx := b - 'a'
		aimNode := cur.next[idx]
		if aimNode == nil {
			aimNode = &Node{}
			cur.next[idx] = aimNode
		}
		cur = aimNode
	}
	cur.isStr = true
}

func (this *Trie) Search(word string) bool {
	var cur *Node = this.root
	for _, b := range word {
		idx := b - 'a'
		aimNode := cur.next[idx]
		if aimNode == nil {
			return false
		}
		cur = aimNode
	}
	return cur.isStr
}

func (this *Trie) StartsWith(prefix string) bool {
	var cur *Node = this.root
	for _, b := range prefix {
		idx := b - 'a'
		aimNode := cur.next[idx]
		if aimNode == nil {
			return false
		}
		cur = aimNode
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
