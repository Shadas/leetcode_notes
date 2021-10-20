package _211_Design_Add_and_Search_Words_Data_Structure

type Node struct {
	isStr bool      // 是否为完整字符串
	next  [26]*Node // a-z 位置可能的后继节点，如果为nil则尚未存在, idx=字符-'a'
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{root: &Node{}}
}

func (this *WordDictionary) AddWord(word string) {
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

func (this *WordDictionary) Search(word string) bool {
	return search(this.root, word)
}

func search(node *Node, word string) bool {
	if len(word) == 0 {
		return node.isStr
	}
	b := word[0]
	if b == '.' {
		for _, n := range node.next {
			if n == nil {
				continue
			}
			ret := search(n, word[1:])
			if ret == true {
				return true
			}
		}
		return false
	} else {
		idx := b - 'a'
		n := node.next[idx]
		if n == nil {
			return false
		}
		return search(n, word[1:])
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
